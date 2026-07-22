package env

import (
	"fmt"
	"strings"

	"github.com/go-errr/go/err"
	"github.com/go-external-config/go/env"
	"github.com/go-jang/go/lang"
	"github.com/go-jang/go/util/optional"
	consul "github.com/hashicorp/consul/api"
)

const CONSUL_KEY_PREFIX = "consul."
const CONSUL_VALUE_PREFIX = "consul:"
const consul_addr = "consul.addr"
const consul_token = "consul.token"

type ConsulPropertySource struct {
	environment *env.Environment
	client      *consul.Client
}

func NewConsulPropertySource() *ConsulPropertySource {
	ps := &ConsulPropertySource{}
	ps.environment = env.Instance()
	ps.client = ps.newClient()
	return ps
}

func (this *ConsulPropertySource) Name() string {
	return "ConsulPropertySource"
}

func (this *ConsulPropertySource) HasProperty(key string) bool {
	if strings.HasPrefix(key, CONSUL_KEY_PREFIX) {
		switch key {
		case consul_addr:
			return false
		case consul_token:
			return false
		default:
			return true
		}
	}
	for _, source := range this.environment.PropertySources() {
		if source.Properties() != nil && source.HasProperty(key) {
			return strings.HasPrefix(source.Property(key), CONSUL_VALUE_PREFIX)
		}
	}
	return false
}

func (this *ConsulPropertySource) Property(key string) string {
	if strings.HasPrefix(key, CONSUL_KEY_PREFIX) {
		return this.getPropertyValue(fmt.Sprint(this.environment.ResolveRequiredPlaceholders(key[len(CONSUL_KEY_PREFIX):])))
	}
	for _, source := range this.environment.PropertySources() {
		if source.Properties() != nil && source.HasProperty(key) {
			return this.getPropertyValue(fmt.Sprint(this.environment.ResolveRequiredPlaceholders(source.Property(key)[len(CONSUL_VALUE_PREFIX):])))
		}
	}
	panic(err.NewIllegalArgumentException("No value present for " + key))
}

func (this *ConsulPropertySource) getPropertyValue(key string) string {
	pair, _, e := this.client.KV().Get(key, nil)
	if e != nil {
		panic(err.NewRuntimeExceptionFrom(fmt.Sprintf("Cannot get Consul property %s", key), e))
	}
	lang.Assert(pair != nil, "No value present for %s", key)
	return string(pair.Value)
}

func (this *ConsulPropertySource) newClient() *consul.Client {
	config := consul.DefaultConfig()
	config.Address = this.environment.Property(consul_addr)
	token := fmt.Sprint(this.environment.ResolveRequiredPlaceholders("${consul.token:}"))
	if len(token) > 0 {
		config.Token = token
	}
	return optional.OfCommaErr(consul.NewClient(config)).OrElsePanic("Unable to initialize Consul client")
}

func (this *ConsulPropertySource) Properties() map[string]string {
	return nil
}

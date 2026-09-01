package consul

import (
	"github.com/go-external-config/consul/env"
	config "github.com/go-external-config/go/env"
)

func init() {
	config.RegisterPropertySource(env.NewConsulPropertySource())
}

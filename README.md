# ConsulPropertySource

HashiCorp Consul is a service networking solution that enables teams to manage secure network connectivity between services and across on-prem and multi-cloud environments and runtimes. Consul offers service discovery, service mesh, traffic management, and automated updates to network infrastructure devices. ([more](https://developer.hashicorp.com/consul/docs/intro))

cmd/app/main.go

    import (
        consul "github.com/go-external-config/consul/env"
        "github.com/go-external-config/go/env"
    )

    var _ = env.Instance().WithPropertySource(consul.NewConsulPropertySource())

    func main() {
    	defer err.Recover()
    	fmt.Println(env.Value[string]("${db.pass}"))
    	// fmt.Println(env.Value[string]("${consul.app/db/password}"))
    }

config/application.yaml

    db:
    	pass: consul:app/db/password

    consul:
    	addr: http://127.0.0.1:8200
    	token: generated

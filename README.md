# ConsulPropertySource

HashiCorp Consul is a service networking solution that enables teams to manage secure network connectivity between services and across on-prem and multi-cloud environments and runtimes. Consul offers service discovery, service mesh, traffic management, and automated updates to network infrastructure devices. ([more](https://developer.hashicorp.com/consul/docs/intro))

cmd/app/main.go

```go
import (
    _ "github.com/go-external-config/consul"
    "github.com/go-external-config/go/env"
)

func main() {
    defer err.Recover()
    fmt.Println(env.Value[string]("${db.pass}"))
    // fmt.Println(env.Value[string]("${consul.app/db/password}"))
}
```

config/application.yaml

```yaml
db:
    pass: consul:app/db/password

consul:
    addr: http://127.0.0.1:8500
    token: generated
```
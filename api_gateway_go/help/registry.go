package help

import (
	"context"

	"github.com/micro/go-micro/registry"

	consul "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func Get_registry() registry.Registry {
	my_registry := registry.NewRegistry(

		func(o *registry.Options) {

			consul_host := viper.GetString("consul.host")
			consul_port := viper.GetString("consul.port")
			consul_url := consul_host + ":" + consul_port
			token := viper.GetString("consul.token")

			consul_config := consul.Config{
				Token:   token,
				Address: consul_url,
			}
			o.Context = context.WithValue(context.TODO(), "consul_config", &consul_config)
		},
	)

	return my_registry

}

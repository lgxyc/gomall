package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/lgxyc/gomall/app/frontend/conf"
	"github.com/lgxyc/gomall/app/frontend/utils"
	frontendutils "github.com/lgxyc/gomall/app/frontend/utils"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	utils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}

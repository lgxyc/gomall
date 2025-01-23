package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/lgxyc/gomall/app/checkout/conf"
	checkoututils "github.com/lgxyc/gomall/app/checkout/utils"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
		initCartClient()
		initPaymentClient()
	})
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(
			transport.GRPC,
		),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoututils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(
			transport.GRPC,
		),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: conf.GetConf().Kitex.Service,
		}),
		client.WithTransportProtocol(
			transport.GRPC,
		),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}

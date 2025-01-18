package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/lgxyc/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/lgxyc/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/lgxyc/gomall/demo/demo_proto/middleware"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	// 客户端集成自定义 middleware
	// client.WithMiddleware
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r), client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware))
	if err != nil {
		log.Fatal(err)
	}

	// withValue 只传递给下个服务,不传递给下下个服务
	// withPersistentValue 可以一直传递下去
	// 将信息放到 metainfo 中
	// meta 传递信息的key风格需要是大写+下划线的CJI风格
	// 例如 CLIENT_NAME
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_thrift")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})

	// 处理异常调用
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v", bizErr)
		}
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}

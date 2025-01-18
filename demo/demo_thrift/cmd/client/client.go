package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/lgxyc/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/lgxyc/gomall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	// 传递rpcinfo
	// 配置metahandler 和 transportProtocol
	cli, err := echo.NewClient("demo_thrift",
		client.WithHostPorts("localhost:8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}))
	if err != nil {
		panic(err)
	}
	req, err := cli.Echo(context.TODO(), &api.Request{
		Message: "hello",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req)
}

package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	wuyuz "github.com/wuyuz/newmicro/proto/wuyuz"
)

type HelloServer struct {
}

func (c *HelloServer) SayHello(ctx context.Context, req *wuyuz.SayRequest, res *wuyuz.SayResponse) error {
	res.Answer = "我们的口号是：\"" + req.Message + "\""
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("wuyuz.wang.server"),
	)

	// 初始化方法
	service.Init()

	//注册服务
	wuyuz.RegisterHelloHandler(service.Server(), new(HelloServer))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

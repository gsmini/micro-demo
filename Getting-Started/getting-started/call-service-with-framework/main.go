package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	proto "github.com/micro/services/helloworld/proto"
)

func main() {
	// 初始化服务
	srv := service.New()

	// 服务之间调用是通过protobuf客户端内部调用的
	client := proto.NewHelloworldService("helloworld", srv.Client())

	// 调用helloworld的服务
	rsp, err := client.Call(context.Background(), &proto.CallRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Message)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}

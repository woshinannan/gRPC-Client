
package main

import (
	"context"
	"fmt"
	"gRPC-Client/private/pb"
	"gRPC-Client/private/service"
)

func main() {
	// 获取客户端
	client := service.GetClient()

	var reqMsg = &pb.ReqMsg{
		Age:                  20,
		Name:                 "haha",
	}

	rspMsg, err := client.Echo(context.Background(), reqMsg)
	if err == nil {
		fmt.Printf("success, reqMsg=%+v, rspMsg=%+v\n", reqMsg, rspMsg)
	} else {
		fmt.Printf("failed, %+v\n", err)
	}
}

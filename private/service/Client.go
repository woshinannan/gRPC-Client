
package service

import (
	"fmt"
	"gRPC-Client/private/pb"
	"google.golang.org/grpc"
)

// 指明服务端的ip和端口
const Address = "127.0.0.1:50052"

func GetClient() pb.MyServiceNameClient {
	// 拨号，创建与gRPC服务端的连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc.Dial failed, err=%+v\n", err)
	}

	// 初始化客户端
	client := pb.NewMyServiceNameClient(conn)

	return client
}
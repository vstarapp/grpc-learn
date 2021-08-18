package main

import (
	"context"
	"fmt"
	"grpc-learn/proto/proto"

	"google.golang.org/grpc"
)

// Address server端地址
const (
	Address string = "localhost:50051"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//连接grpc
	client := proto.NewSpiderClient(conn)
	req := proto.SpiderRequest{
		Address: "http://www.baidu.com",
		Method:  "GET",
	}
	// 调用server的注册方法
	resp, err := client.GetAddressResponse(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印返回值
	fmt.Println(resp)
}

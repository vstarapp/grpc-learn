// server 端
package main

import (
	"context"
	"grpc-learn/proto/proto"
	"io/ioutil"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

const (
	// Address 监听地址
	Address string = "localhost:50051"
	// Method 通信方法
	Method string = "tcp"
)

func (s *server) GetAddressResponse(ctx context.Context, a *proto.SpiderRequest) (*proto.SpiderResponse, error) {
	switch a.Method {
	case "get", "Get", "GET":
		status, body, err := get(a.Address)
		if err != nil {
			return nil, err
		}
		res := proto.SpiderResponse{
			HttpCode: int32(status),
			Response: body,
		}

		return &res, err
	}

	return nil, nil
}

func get(address string) (s int, r string, err error) {
	resp, err := http.Get(address)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	s = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r = string(body)
	return
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Method, Address)
	if err != nil {
		return
	}
	s := grpc.NewServer()                    // 创建GRPC
	proto.RegisterSpiderServer(s, &server{}) // 在GRPC服务端注册服务

	reflection.Register(s) // 在GRPC服务器注册服务器反射服务
	// Serve方法接收监听的端口,每到一个连接创建一个ServerTransport和server的grroutine
	// 这个goroutine读取GRPC请求,调用已注册的处理程序进行响应
	err = s.Serve(listener)
	if err != nil {
		return
	}
}

# grpc-learn
一个高性能、通用的开源RPC框架，其由Google主要面向移动应用开发并基于HTTP/2协议标准而设计，基于ProtoBuf(Protocol Buffers)序列化协议开发，且支持众多开发语言。gRPC基于HTTP/2标准设计，带来诸如双向流控、头部压缩、单TCP连接上的多复用请求等特性。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。
学习笔记：
创建项目目录
mkdir -p grpc-learn && cd grpc-learn
使用mod初始化
go mod init
安装相关包#
安装 golang 的proto工具包
go get -u github.com/golang/protobuf/proto
安装 goalng 的proto编译支持
go get -u github.com/golang/protobuf/protoc-gen-go
安装 GRPC 包
go get -u google.golang.org/grpc

生成 .bp.go 文件
输出的目录 proto所在目录
protoc --go_out=plugins=grpc:. proto/spider.proto
运行后会在当前目录下生成 spider.pb.go 文件

运行#
需要先启动 server 端监听端口,再启动 client 端向端口发送请求
我们运行后可看到结果已经正常返回并打印
go run .\server\main.go
go run .\client\client.go

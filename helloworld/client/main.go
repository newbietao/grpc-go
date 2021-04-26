package main

import (
	"context"
	"fmt"
	"time"

	pd "github.com/newbietao/grpc-go/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 创建一个连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 初始化一个客户端
	client := pd.NewHelloWorldClient(conn)
	// 创建ctx
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用grpc
	res, err := client.SayHello(ctx, &pd.HelloReq{Name: "world"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

package main

import (
	"context"
	"fmt"
	"net"

	pd "github.com/newbietao/grpc-go/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// 实现service 接口
type service struct{}

func (s *service) SayHello(ctx context.Context, req *pd.HelloReq) (res *pd.HelloRes, err error) {
	return &pd.HelloRes{Msg: "hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pd.RegisterHelloWorldServer(s, &service{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}

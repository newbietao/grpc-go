package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/newbietao/grpc-go/grpc-user/user"
	"google.golang.org/grpc"
)

const (
	addr = "127.0.0.1:5555"
)

func main() {
	// 创建一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("grpc.Dial:", err)
		return
	}
	defer conn.Close()
	// 创建一个客户端
	client := user.NewUserServiceClient(conn)
	// 添加用户
	for i := 0; i < 3; i++ {
		u := user.User{
			Name:   "name" + strconv.Itoa(i+1),
			Age:    int32(20 + i),
			Hight:  int32(180 + i),
			Gender: user.User_Female,
		}
		fmt.Println(u)
		res, err := client.AddUser(context.Background(), &u)
		fmt.Println("AddUser:", res, err)
	}
	// 获取用户
	res, err := client.GetUserByName(context.Background(), &user.GetUserReq{Name: "name1"})
	fmt.Println("GetUserByName:", res, err)
	// 删除用户
	res2, err := client.DelUserByName(context.Background(), &user.DelUserReq{Name: "name1"})
	fmt.Println("DelUserByName:", res2, err)
	// 获取用户
	res3, err := client.GetUserByName(context.Background(), &user.GetUserReq{Name: "name1"})
	fmt.Println("GetUserByName:", res3, err)
}

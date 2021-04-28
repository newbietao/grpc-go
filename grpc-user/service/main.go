package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/newbietao/grpc-go/grpc-user/user"
	"google.golang.org/grpc"
)

var UserDB = make(map[string]*user.User, 100)

type UserService struct {
	lock sync.Mutex
}

func (u *UserService) GetUserByName(ctx context.Context, req *user.GetUserReq) (res *user.User, err error) {
	res = &user.User{}
	name := req.Name
	if name == "" {
		return res, errors.New("name is empty!")
	}
	u.lock.Lock()
	defer u.lock.Unlock()
	if user, ok := UserDB[name]; ok {
		return user, nil
	}

	return res, errors.New("user not exist!")
}

func (u *UserService) AddUser(ctx context.Context, userInfo *user.User) (res *user.CommonUserRes, err error) {
	res = &user.CommonUserRes{}
	if userInfo.Name == "" {
		err = errors.New("参数错误")
		res.Code = 1
		res.Msg = err.Error()
		fmt.Println(res)
		return
	}
	u.lock.Lock()
	UserDB[userInfo.Name] = userInfo
	u.lock.Unlock()
	fmt.Println(res)
	return
}
func (u *UserService) DelUserByName(ctx context.Context, req *user.DelUserReq) (res *user.CommonUserRes, err error) {
	res = &user.CommonUserRes{}
	u.lock.Lock()
	defer u.lock.Unlock()
	if _, ok := UserDB[req.Name]; ok {
		delete(UserDB, req.Name)
		res.Msg = "删除成功"
		return
	}
	res.Code = 1
	res.Msg = "删除失败"
	return
}

const (
	addr = "127.0.0.1:5555"
)

func main() {
	us := &UserService{}
	// 创建一个监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("net.Listen: ", err)
		return
	}
	// 实例化Grpc服务
	server := grpc.NewServer()
	// 注册服务
	user.RegisterUserServiceServer(server, us)
	// 启动服务
	server.Serve(lis)
}

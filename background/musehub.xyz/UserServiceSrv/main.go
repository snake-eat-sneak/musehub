// Package main
package main

import (
	proto "UserServiceSRV/proto"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

type User struct{}

func (s *User) UserLogin(ctx context.Context, req *proto.UserLoginRequest, rsp *proto.UserLoginResponse) error {
	log.Log("Received Say.Hello request")
	userInfo := make([]*proto.UserInfo,0,1)
	userInfo = append(userInfo, &proto.UserInfo{UserId:req.UserId, Password:req.Password, Name:"ctwww_test"})
	rsp.Code = 0
	rsp.UserInfo = userInfo
	return nil
}

func main() {
	//etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(
		micro.Name("srv.musehub.xyz.user"),
		micro.Address(":8011"),
		//micro.Registry(etcdReg),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	proto.RegisterUserHandler(service.Server(), new(User))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
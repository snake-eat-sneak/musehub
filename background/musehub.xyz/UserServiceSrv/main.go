// Package main
package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"

	proto "UserServiceSrv/proto"
)

type User struct{}

func (s *User) UserLogin(ctx context.Context, req *proto.UserLoginRequest, rsp *proto.UserLoginResponse) error {
	log.Log("Received Say.Hello request")
	userInfo := make([]*proto.User,1,1)
	userInfo = append(userInfo, &proto.User{UserId:req.UserId, Password:req.Password, Name:"ctwww_test"})
	rsp.Code = 0
	rsp.User = userInfo
	return nil
}

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("srv.musehub.xyz.user"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	proto.RegisterUserServiceSrvHandler(service.Server(), new(User))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
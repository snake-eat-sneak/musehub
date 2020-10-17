package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"

	proto "UserServiceAPI/proto"
	user "musehub.xyz/UserServiceSrv/proto"
	//"github.com/micro/go-micro/v2"

	"context"
)

type User struct {
	Client user.UserServiceSrvService
}

func (u *User) UserLogin(ctx context.Context, req *proto.UserLoginRequest, rsp *proto.UserLoginResponse) error {
	log.Print("Received User.UserLogin API request")

	// make the request
	response, err := u.Client.UserLogin(ctx, &user.UserLoginRequest{UserId: req.UserId, Password:req.Password})
	if err != nil {
		return err
	}

	// set api response
	ResUser := make([]*proto.User, 0, 1)
	if len(response.User) > 0{
		userInfo := *response.User[0]
		ResUser = append(ResUser, &proto.User{UserId: userInfo.UserId, Password: userInfo.Password, Name:userInfo.Name})
	}
	rsp.User = ResUser
	rsp.Code = response.Code
	return nil
}

func main() {
	// Create service
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("api.musehub.xyz.user"),
		micro.Registry(consulReg),
		micro.Address(":8021"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	proto.RegisterUserServiceAPIHandler(service.Server(), &User{
		Client: user.NewUserServiceSrvService("go.micro.srv.user", service.Client()),
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
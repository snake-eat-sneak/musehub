package main

import (
	"github.com/micro/go-micro"
	"log"

	proto "UserServiceAPI/proto"
	user "musehub.xyz/UserServiceSRV/proto"
	//"github.com/micro/go-micro/v2"

	"context"
)

type User struct {
	Client user.UserService
}

func (u *User) UserLogin(ctx context.Context, req *proto.UserLoginRequest, rsp *proto.UserLoginResponse) error {
	log.Print("Received User.UserLogin API request")

	// make the request
	response, err := u.Client.UserLogin(ctx, &user.UserLoginRequest{UserId: req.UserId, Password:req.Password})
	if err != nil {
		return err
	}

	// set api response
	ResUser := make([]*proto.UserInfo, 0, 1)
	if len(response.UserInfo) > 0{
		userInfo := *response.UserInfo[0]
		ResUser = append(ResUser, &proto.UserInfo{UserId: userInfo.UserId, Password: userInfo.Password, Name:userInfo.Name})
	}
	rsp.UserInfo = ResUser
	rsp.Code = response.Code
	return nil
}

func main() {
	// Create service
	//etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(
		micro.Name("api.musehub.xyz.user"),
		//micro.Registry(etcdReg),
		micro.Address(":8021"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	proto.RegisterUserHandler(service.Server(), &User{
		Client: user.NewUserService("srv.musehub.xyz.user", service.Client()),
	})

	// for handler use
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
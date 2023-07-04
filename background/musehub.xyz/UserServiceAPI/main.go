package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"
	"log"

	api "musehub.xyz/proto/api"
	user "musehub.xyz/proto/srv"
	//"github.com/micro/go-micro/v2"

	"context"
)

type User struct {
	Client user.UserService
}

func (u *User) UserLogin(ctx context.Context, req *api.UserLoginRequest, rsp *api.UserLoginResponse) error {
	log.Print("Received User.UserLogin API request")

	etcd.NewRegistry()

	// make the request
	response, err := u.Client.UserLogin(ctx, &user.UserLoginRequest{UserId: req.UserId, Password:req.Password})
	if err != nil {
		return err
	}

	// set api response
	ResUser := make([]*api.UserInfo, 0, 1)
	if len(response.UserInfo) > 0{
		userInfo := *response.UserInfo[0]
		ResUser = append(ResUser, &api.UserInfo{UserId: userInfo.UserId, Password: userInfo.Password, Name:userInfo.Name})
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
	api.RegisterUserHandler(service.Server(), &User{
		Client: user.NewUserService("srv.musehub.xyz.user", service.Client()),
	})

	// for handler use
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
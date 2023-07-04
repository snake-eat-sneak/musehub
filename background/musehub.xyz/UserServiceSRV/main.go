// Package main
package main

import (
	srv "musehub.xyz/proto/srv"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

type User struct{}

func (s *User) UserLogin(ctx context.Context, req *srv.UserLoginRequest, rsp *srv.UserLoginResponse) error {
	log.Log("Received Say.Hello request")
	userInfo := make([]*srv.UserInfo,0,1)
	userInfo = append(userInfo, &srv.UserInfo{UserId:req.UserId, Password:req.Password, Name:"ctwww_test"})
	rsp.Code = 0
	rsp.UserInfo = userInfo
	return nil
}

func main() {
	//etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	service := micro.NewService(
		micro.Name("srv.musehub.xyz.user"),
		//micro.Address(":8011"),
		//micro.Registry(etcdReg),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	srv.RegisterUserHandler(service.Server(), new(User))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
package main

import (
	"fmt"
	"gateway/example"
	"gateway/helper"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	fmt.Println("start===================")

	// 路由
	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	v1Group.Handle("GET", "/user", func(context *gin.Context) {
		fmt.Println("user========")
		context.String(200, "user api")
	})
	v1Group.Handle("GET", "/news", func(context *gin.Context) {
		fmt.Println("news========")
		context.String(200, "news api")
	})
	v1Group.Handle("POST", "/example", func(context *gin.Context) {

		var pr helper.GatewayRequest
		err := context.Bind(&pr)
		if err != nil{
			pr = helper.GatewayRequest{Name:"null"}
		}
		res := gin.H{"data": example.GetExampleContent(pr.Name)}
		context.JSON(200, res)
	})

	// 服务注册
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	server := web.NewService(
		web.Name("musehub_gateway"),
		//web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)

	server.Init()
	server.Run()
}

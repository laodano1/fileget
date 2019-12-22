package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"

	"log"
)

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.146:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),          // 注册服务到consul
		micro.Name("helloWorld-srv"), // 真正的服务名称
	)

	service.Init()

	err := service.Run()
	if err != nil {
		log.Fatalf("server run error: %v", err)
	}

	log.Printf("server name: %v", service.Server().String())

}

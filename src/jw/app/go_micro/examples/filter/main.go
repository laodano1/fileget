package main

import (
	"context"
	"fmt"

	"github.com/micro/examples/filter/version"
	proto "github.com/micro/examples/service/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService()
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(
		// provide a context
		context.TODO(),
		// provide the request
		&proto.HelloRequest{Name: "John"},
		// set the filter
		version.Filter("latest"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}

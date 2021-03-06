package main

import (
	"fmt"
	"log"

	"context"
	proto "github.com/micro/examples/service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		// wrap the handler
		micro.WrapHandler(logWrapper),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

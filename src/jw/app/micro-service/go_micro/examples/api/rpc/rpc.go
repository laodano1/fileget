package main

import (
	"log"

	proto "github.com/micro/examples/api/rpc/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"

	"context"
)

type Example struct{}

type Foo struct{}

// Example.Call is a method which will be served by http request /example/call
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /example/call goes to go.micro.api.example Example.Call
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Received Example.Call request")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	rsp.Message = "got your request " + req.Name
	return nil
}

// Foo.Bar is a method which will be served by http request /example/foo/bar
// Because Foo is not the same as the service name it is mapped beyond /example/
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Received Foo.Bar request")

	// noop

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// register example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// register foo handler
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

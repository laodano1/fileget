package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	//api "github.com/micro/micro/api/proto"
)

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.146:8500",
		}
	})

	meta := make(map[string]string)
	meta["host"] = "192.168.1.146"
	meta["port"] = "9999"

	service := micro.NewService(
		micro.Registry(reg),     // 注册服务到consul
		micro.Name("hwService"), // 真正的服务名称
		micro.Metadata(meta),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
	)

	service.Server().Init(
		server.Id("jw-test-server"),
	)

	//service.Server().Handle(
	//	service.Server().NewHandler(new(Redirect)),
	//)

	err := service.Run()
	if err != nil {
		log.Fatalf("server run error: %v", err)
	}

}

//type Redirect struct{}
//
//func (r *Redirect) Url(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	rsp.StatusCode = int32(301)
//	rsp.Header = map[string]*api.Pair{
//		"Location": &api.Pair{
//			Key:    "Location",
//			Values: []string{"https://google.com"},
//		},
//	}
//	return nil
//}

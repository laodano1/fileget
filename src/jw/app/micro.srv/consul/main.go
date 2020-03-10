package main

import (
	"github.com/davyxu/golog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"time"
)

func main() {
	meta := make(map[string]string)
	//meta[""]
	logger := golog.New("tttt")
	svc := micro.NewService(
		micro.Name("game-600101"),
		micro.Registry(consul.NewRegistry(
			registry.Addrs(
				"10.0.0.119:8500",
			))),
		micro.RegisterTTL(10 * time.Second),
		micro.Address(":8801"),
		micro.Metadata(meta),
		)


     if err := svc.Run(); err != nil {
     	logger.Errorf("service run failed: %v", err)
	 }
}

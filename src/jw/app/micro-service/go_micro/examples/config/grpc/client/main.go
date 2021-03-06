package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	grpcConfig "github.com/micro/go-plugins/config/source/grpc"
)

type Micro struct {
	Info
}

type Info struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
}

func main() {
	// create new source
	source := grpcConfig.NewSource(
		grpcConfig.WithAddress("127.0.0.1:8600"),
		grpcConfig.WithPath("/micro"),
	)

	// create new config
	conf := config.NewConfig()

	// load the source into config
	if err := conf.Load(source); err != nil {
		log.Fatal(err)
	}

	configs := &Micro{}
	if err := conf.Scan(configs); err != nil {
		log.Fatal(err)
	}

	log.Logf("Read config: %s", string(conf.Bytes()))

	// watch the config for changes
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watching for changes ...")

	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watching for changes: %v", string(v.Bytes()))
	}
}

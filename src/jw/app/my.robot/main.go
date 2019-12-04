package main

import (
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
	"log"
	"runtime"

	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

var (
	srvAdd  = ":8888"
	logger  *golog.Logger
	memStat *runtime.MemStats
)

var (
	consulHost = "192.168.1.146:8500"
)

func main() {
	logger = golog.New("my.robot")
	logger.SetParts()

	config := api.DefaultConfig()
	config.Address = consulHost

	cli, err := api.NewClient(config)
	if err != nil {
		log.Println("new api client error:", err)
		return
	}

	as, err := cli.Agent().Services()
	if err != nil {
		log.Println("Get Services error:", err)
		return
	}

	for name, s := range as {
		log.Printf("service name: %s, meta: %v, service: %v\n", name, s.Meta, s.Service)
	}

}

package main

import (
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
)

var logger = golog.New("service.discovery-2")

func main() {

	logger.Infof("New consul client")
	config := api.DefaultConfig()
	config.Address = "192.168.1.146:8500"
	client, err := api.NewClient(config)
	if err != nil {
		logger.Errorf("new client error: %v", err)
	}

	services, metainfo, err := client.Health().Service("jw.health-1", "", true, &api.QueryOptions{
		//Datacenter:        "",
		//AllowStale:        false,
		//RequireConsistent: false,
		//UseCache:          false,
		//MaxAge:            0,
		//StaleIfError:      0,
		//WaitIndex:         0,
		//WaitHash:          "",
		//WaitTime:          0,
		//Token:             "",
		//Near:              "",
		//NodeMeta:          nil,
		//RelayFactor:       0,
		//LocalOnly:         false,
		//Connect:           false,
		//Filter:            "",
	})

	logger.Infof("service: %v, metainfo: %v", services, metainfo)

	for _, v := range services {
		logger.Infof("srv: %v, %v:%v, %v", v.Service.Service, v.Service.Address, v.Service.Port, v.Service.TaggedAddresses)
	}

	//for _, v := range metainfo {
	//
	//}

}

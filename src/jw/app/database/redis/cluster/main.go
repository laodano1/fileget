package main

import (
	"fileget/util"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	clusterCli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.0.0.130:6379",
			"10.0.0.130:6381",
			"10.0.0.130:6382",
			"10.0.0.130:6383",
			"10.0.0.130:6384",
			"10.0.0.130:6385",
		},
	})

	util.Lg.Debugf("client id: %v", clusterCli.ClusterSlots())
	util.Lg.Debugf("%v", clusterCli.Ping())

	util.Lg.Debugf("set status: %v", clusterCli.SetNX("keys11", 11, 10*time.Minute))
}

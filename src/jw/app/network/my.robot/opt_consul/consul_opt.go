package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
	"net/http"
	"time"
)

var (
	//consulHost = "http://192.168.11.79:8500"
	consulHost = "http://10.0.0.34:8500"
	logger = golog.New("consul-operation")

)


type PutValue struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data  []byte  `json:"data"`
}

type ConsuOpt struct {
	ConsCli *api.Client
}

func (co *ConsuOpt) CreateKey() {
	k := fmt.Sprintf("test/agent-%v", + time.Now().UnixNano())

	pv := PutValue{
		Code:    111,
		Message: "Hello World",
		Data:    []byte("test1: 1, test2: \"2\""),
	}

	pvJson, err := json.MarshalIndent(pv, "", "    ")
	if err != nil {
		logger.Errorf("json marshal error: %v", err)
	}

	v := bytes.NewBuffer(pvJson)

	urlStr := consulHost + "/v1/kv/" + k

	req, err := http.NewRequest("PUT", urlStr, v)
	if err != nil {
		logger.Errorf("new http request error: %v", err)
	}

	httpCli := http.Client{ Timeout: 5 * time.Second}
	resp, err := httpCli.Do(req)
	if err != nil {
		logger.Errorf("do http request error: %v", err)
	}

	logger.Infof("get server status: '%v'", resp.Status)
}

func (co *ConsuOpt) GetKeys() {

	kvs, _, _ := co.ConsCli.KV().List("test", &api.QueryOptions{})
	for _, v := range kvs {
		tpv := &PutValue{}
		json.Unmarshal(v.Value, tpv)
		logger.Infof("key: %v, value: %v", v.Key, tpv.Code)
	}
}

func (co *ConsuOpt) DelKeys() {
	kvs, _, _ := co.ConsCli.KV().List("test", &api.QueryOptions{})
	for _, v := range kvs {
		co.ConsCli.KV().Delete(v.Key, &api.WriteOptions{})
		tpv := &PutValue{}
		json.Unmarshal(v.Value, tpv)
		logger.Infof("delete key: %v, value: %v", v.Key, tpv.Code)
	}
}

func NewConOpt() *ConsuOpt {
	consuCfg := api.DefaultConfig()
	consuCfg.Address = consulHost
	conCli, err := api.NewClient(consuCfg)
	if err != nil {
		logger.Errorf("new consul client error: %v", err)
	}
	return &ConsuOpt{ConsCli: conCli}
}

func (co *ConsuOpt) GetAndDelDeadServies() {

	checks, err := co.ConsCli.Agent().Checks()
	if err != nil {
		logger.Debugf("client.Agent().Checks() error: %v\n", err)
		panic("client.Agent().Checks() error: %v\n")
	}
	var okServices string
	for name, ac := range checks {
		//if ac.ServiceName == "wanke-gamesvr" && ac.Status == api.HealthCritical {
		//if ac.ServiceName == "game-600101" && ac.Status == api.HealthCritical {
		if ac.ServiceName == "game-600102" && ac.Status == api.HealthCritical {
			//log.Debugf("name: %27v, agent check: %-10v\n", name, ac.Status)
			okServices = name + " " + okServices
			logger.Debugf("service name: %v", name)
			err := co.ConsCli.Agent().ServiceDeregister(name)
			if err != nil {
				logger.Debugf("deregister service '%v' failed1", name)
				panic(err)
			}
		}
	}


}

func main() {
	logger.EnableColor(true)
	logger.SetColor("red")

	co := NewConOpt()
	//co.CreateKey()
	//co.GetKeys()
	//co.DelKeys()
	co.GetAndDelDeadServies()

}

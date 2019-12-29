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
	consulHost = "http://10.0.0.143:8500"
	logger = golog.New("consul-operation")

)


type PutValue struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data  []byte  `json:"data"`
}


func CreateKey() {
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

func GetKeys() {
	consuCfg := api.DefaultConfig()
	consuCfg.Address = consulHost
	conCli, err := api.NewClient(consuCfg)
	if err != nil {
		logger.Errorf("new consul client error: %v", err)
	}

	kvs, _, err := conCli.KV().List("test", &api.QueryOptions{})
	for _, v := range kvs {
		tpv := &PutValue{}
		json.Unmarshal(v.Value, tpv)
		logger.Infof("key: %v, value: %v", v.Key, tpv.Code)
	}
}

func DelKeys() {
	consuCfg := api.DefaultConfig()
	consuCfg.Address = consulHost
	conCli, err := api.NewClient(consuCfg)
	if err != nil {
		logger.Errorf("new consul client error: %v", err)
	}

	kvs, _, err := conCli.KV().List("test", &api.QueryOptions{})
	for _, v := range kvs {
		conCli.KV().Delete(v.Key, &api.WriteOptions{})
		tpv := &PutValue{}
		json.Unmarshal(v.Value, tpv)
		logger.Infof("delete key: %v, value: %v", v.Key, tpv.Code)
	}
}

func main() {
	logger.EnableColor(true)
	logger.SetColor("red")
	CreateKey()
	GetKeys()
	//DelKeys()
}

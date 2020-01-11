package main

import (
	"encoding/json"
	"github.com/davyxu/golog"
	"net"
	"os"
	"strings"
)

// 启动配置模板
var InitConf = map[string]string{
	"agent": `{
  "nodeConfig": [
    {
      "host": "HOSTIP:8801",
      "type": "http.Acceptor",
      "name":"game-agent",
      "protoType": "http",
      "description": ""
    }
  ]
}`,
	"match": `{
  "nodeConfig": [
    {
      "host": "HOSTIP:8900",
      "type": "tcp.Acceptor",
      "name":"match",
      "protoType": "tcp.ltv",
      "description": ""
    }
  ]
}`,
	"doudizhu": `{
  "nodeConfig": [
    {
      "host": "HOSTIP:8813",
      "type": "tcp.Acceptor",
      "name":"doudizhu",
      "protoType": "tcp.ltv",
      "description": ""
    }
  ]
}`,
	"xuezhandaodi": `{
  "nodeConfig": [
    {
      "host": "HOSTIP:8803",
      "type": "tcp.Acceptor",
      "name":"xuezhandaodi",
      "protoType": "tcp.ltv",
      "description": ""
    }
  ]
}`,

	"mysql":  `{
      "host": "MYSQLUSER:MYSQLPWD@tcp(MYSQLHOSTIP:MYSQLPORT)/game?charset=utf8&parseTime=True&loc=Local",
      "name": "mysql"
}`,
	"redis":    `{
      "host": "HOSTIP:6379",
      "name": "redis"
}`,
}

var logger = golog.New("json-test")

//连接配置
type NodeConfig struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ProtoType            string   `protobuf:"bytes,4,opt,name=protoType,proto3" json:"protoType,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	MetaNode             bool     `protobuf:"varint,7,opt,name=metaNode,proto3" json:"metaNode,omitempty"`
}

//服务器连接配置
type ServerConfig struct {
	NodeConfig           []*NodeConfig `protobuf:"bytes,2,rep,name=nodeConfig,proto3" json:"nodeConfig,omitempty"`
}

// load local config from template
func LoadServiceConfigLocal(srvName string) (config *ServerConfig, err error)  {
	config = &ServerConfig{}
	err = LoadJsonLocal(srvName, config)
	//config
	return
}

func LoadJsonLocal( srvName string, val interface{}) (err error) {
	switch srvName {
	case "agent":
		str := InitConf[srvName]
		if err := json.Unmarshal([]byte(str), val); err != nil {
			return err
		}
		logger.Infof("agent config ok")

	case "match":

	case "doudizhu":

	case "xuezhandaodi":

		//case "":

	}

	return
}

func main() {
	config, err := LoadServiceConfigLocal("agent")
	if err != nil {
		panic(err)
	}

	nc := config.NodeConfig[0]

	host, _ := os.Hostname()
	logger.Infof("1. host: %v", host)

	addrs, _ := net.LookupIP(host)
		for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil && !ipv4.IsGlobalUnicast(){
			logger.Infof("IPv4: '%v'", ipv4)
		}
	}

	nc.Host = strings.ReplaceAll(nc.Host, "HOSTIP", "8.8.8.9")
	logger.Infof("2. host: %v", nc)

}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"time"
)

var (
	lg = golog.New("my.nats.publisher")
)

type Message struct {
	Header map[string]string `json:"header"`
	Body   []byte           `json:"body"`
}

// platform <- svc
type NatsMessage struct {
	WsId                 string   `protobuf:"bytes,1,opt,name=wsId,proto3" json:"wsId,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	WsIds                []string `protobuf:"bytes,3,rep,name=wsIds,proto3" json:"wsIds,omitempty"`
}

func main() {
	nc, err := nats.Connect("nats://10.0.0.251:4222", nats.Name("API PublishBytes Example"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		lg.Infof("drain all msg!")
		if err := nc.Drain(); err != nil {
			log.Fatal(err)
		}
		nc.Close()
	}()

	tc := time.Tick(3 * time.Second)
	for {
		select {
		case <-tc:
			pkg := &ServiceStatusACK{
			Result: int32(http.StatusAccepted),
			Error:  "服务器维护中",
			GameId: "",
		}
			pkgj, _ := json.Marshal(pkg)
			encoded, _ := Encode(http.StatusOK, "", pkgj)
			body := &NatsMessage{
				WsId: "1",
				Data:  encoded,
			}
			lg.Infof("publish info: %v", body)
			pi, _ := json.Marshal(body)
			data := &Message{Body: pi}
			pubInfo, _ := json.Marshal(data)
			topic := fmt.Sprintf("wsgw.message.service.%v",  "gateway_http-10.0.0.37-9087")
			if err := nc.Publish(topic, pubInfo); err != nil {
				log.Fatal(err)
			}
		}
	}
}

type ServiceStatusACK struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	GameId               string   `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
}

type GatewayProtoBase  = PostResult

type PostResult struct {
	Code    int    `json:"code"`    //200:成功 详见【状态码】
	Message string `json:"message"` //接口描述信息
	Data    []byte `json:"data"`    //请求json数据
}

func Encode(code int, message string, raw []byte) (bytes []byte, err error) {
	base := &GatewayProtoBase{
		Code:    code,
		Message: message,
		Data:    raw,
	}
	bytes, err = json.Marshal(base)
	return
}
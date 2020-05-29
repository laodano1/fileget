package main

import (
	"context"
	"fileget/src/jw/app/micro.srv/grpc-tcp/proto/myserver"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/codec/grpc"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/tcp"
	"strings"
	"sync/atomic"
	"time"
)

var (
	Host = "10.0.0.8:7777"
	cs   = "10.0.0.32:8500"
	lg   = golog.New("grpc-client")
	cli  client.Client
)

func main() {
	CallRPC()
}

func CallRPC() {

	var c client.Client
	if c == nil {
		lg.Debugf("create a new go micro client! | %v", time.Now().Format(time.RFC3339Nano))
		c = client.NewClient(

			client.Transport(tcp.NewTransport()),
			client.Registry(consul.NewRegistry(registry.Addrs(cs))),
			//client.Codec("application/protobuf", proto.NewCodec),
			//client.Codec("application/grpc", grpc.NewCodec),
			client.Codec("application/grpc+json", grpc.NewCodec),
			//client.Codec("application/octet-stream", raw.NewCodec),
			//client.Codec("application/json-rpc", jsonrpc.NewCodec),
		//client.WrapCall(wrap),
		//client.Retries(3), // retry times
		)
	}

	mysCli := myserver.NewMyServerService("grpc-tcp", c)

	// calling remote address
	//wg := &sync.WaitGroup{}
	cnt := 100
	//wg.Add(cnt)
	var fromCnt int32
	sigCh := make(chan bool, 100)
	for i := 0; i < cnt; i++ {
		go func(id int) {
			msgRsp, err := mysCli.Message(context.Background(), &myserver.MessageReq{
				UserId:  0,
				WsId:    "",
				Service: "service1",
				Token:   "token1",
				//}, client.WithAddress(Host))
			})
			if err != nil {
				lg.Errorf("Call Message interface failed: %v", err)
			} else {
				lg.Debugf("rpc response(%2v): %v", id, msgRsp)
				if strings.Contains(msgRsp.Host, "156:77") {
					fromCnt = atomic.AddInt32(&fromCnt, 1)
					lg.Debugf("====================== from cnt: %v", atomic.LoadInt32(&fromCnt))
				}
			}
			sigCh <- true
			//wg.Done()
		}(i)
	}
	//wg.Wait()
	//time.Sleep(2100 * time.Microsecond)
	var i int
	for {
		select {
		case <-sigCh:
			i++
			if i == cnt {
				lg.Debugf("bye bye")
				return
			}
		}
	}

}

// calling remote address
func Invoke(c client.Client, endpoint string, id int) {
	var rsp interface{}
	req := c.NewRequest("grpc-tcp", endpoint, nil)
	switch endpoint {
	case "MyServer.ConnectA":
		rsp = &myserver.ConnectRsp{}
	case "MyServer.Disconnect":
		rsp = &myserver.DisconnectRsp{}
	case "MyServer.Message":
		rsp = &myserver.MessageRsp{}
	default:
		lg.Debugf("unknown endpoint!")
		return
	}
	//if err := c.Call(context.Background(), req, rsp, client.WithAddress(Host)); err != nil {
	if err := c.Call(context.Background(), req, rsp); err != nil {
		lg.Errorf("call(%2v) with error： %v", id, err)
		return
	}

}

package main

import (
	"context"
	"fileget/src/jw/app/micro.srv/grpc-tcp/proto/myserver"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/tcp"
)



var (
	Host   = "10.0.0.20:7777"
	cs = "10.0.0.60:8500"
    lg = golog.New("grpc-client")
    cli client.Client
)
//
//func Message(grpcSrv string, endpoint string, req *myserver.MessageReq) (err error) {
//	lg.EnableColor(true)
//	//opt1 := client.Registry(consul.NewRegistry(registry.Addrs("localhost:8500")))
//	//opt2 := client.ContentType("application/json") //.Codec("application/json", codec.NewCodec)
//	//client.Codec("application/json", codec.NewCodec)
//	if cli == nil {
//		cli = client.NewClient()
//	}
//
//	//request := c.NewRequest("wanke.platform.api.gate", "Websocket.Message", req)
//	request := cli.NewRequest("wanke.platform.api.gate", endpoint, req)
//
//	var rsp json.RawMessage
//	if err = cli.Call(context.Background(), request, &rsp, client.WithAddress(grpcSrv),); err != nil {
//		err = errors.New(fmt.Sprintf("go micro grpc call failed: %v", err))
//	}
//	return
//}

func main() {
	//lg.SetParts()
	//mreq := &myserver.MessageReq{
	//	UserId:               9876,
	//	WsId:                 "111",
	//	Service:              "srv-112233",
	//	Token:                "aaa",
	//}
	//err := Message(Host, "Websocket.Message", mreq)
	//if err != nil {
	//	lg.Errorf("invoke grpc message interface failed: %v", err)
	//	return
	//}
	//lg.Debugf("invoke grpc message successfully!")
	CallRPC()
}

func CallRPC()  {

	var c client.Client
	if c == nil {
		c = client.NewClient(
			client.Transport(tcp.NewTransport()),
			client.Registry(consul.NewRegistry(registry.Addrs(cs))),
			//client.WrapCall(wrap),
		)
	}



	// calling remote address
	//wg := &sync.WaitGroup{}
	cnt := 10
	//wg.Add(cnt)
	sigCh := make(chan bool, 100)
	for i := 0; i < cnt; i++ {
		go func() {
			Invoke(c, "MyServer.ConnectA")
			Invoke(c, "MyServer.Disconnect")
			Invoke(c, "MyServer.Message")
			sigCh <- true
			//wg.Done()
		}()
	}
	//wg.Wait()
	//time.Sleep(2100 * time.Microsecond)
	var i int
	for {
		select {
		case <- sigCh :
			i++
			if i == cnt {
				lg.Debugf("bye bye")
				return
			}
		}
	}

}

// calling remote address
func Invoke(c client.Client, endpoint string)  {
	var rsp interface{}
	req := c.NewRequest("grpc-tcp", endpoint, nil)
	switch endpoint {
	case "MyServer.ConnectA" :
		rsp = &myserver.ConnectRsp{}
	case "MyServer.Disconnect" :
		rsp = &myserver.DisconnectRsp{}
	case "MyServer.Message" :
		rsp = &myserver.MessageRsp{}
	default:
		lg.Debugf("unknown endpoint!")
		return
	}
	if err := c.Call(context.Background(), req, rsp, client.WithAddress(Host)); err != nil {
		lg.Errorf("call with error： %v", err)
		return
	}
	lg.Debugf("rpc response: %v", rsp)
}
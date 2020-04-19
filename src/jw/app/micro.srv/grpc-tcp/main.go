package main

import (
	"context"
	"fileget/src/jw/app/micro.srv/grpc-tcp/proto/myserver"
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

var (
	logger = golog.New("grpc-tcp")
	PORT   = "10.0.0.156:7778"
	cs     = "10.0.0.131:8500"
)

type mySvr struct {
	cnt int
}

func (ms *mySvr) ConnectA(ctx context.Context, req *myserver.ConnectReq, rsp *myserver.ConnectRsp) (err error) {
	ms.cnt++
	logger.Debugf("in ConnectA => req(%2v) user id: %v", ms.cnt, req.UserId)
	rsp.Confirm = true
	rsp.Host = PORT

	return
}

// 用户断线
func (ms *mySvr) Disconnect(ctx context.Context, dreq *myserver.DisconnectReq, drsp *myserver.DisconnectRsp) (err error) {
	logger.Debugf("in Disconnect => dreq user id: %v", dreq.UserId)

	drsp.Confirm = true
	drsp.Host = PORT
	return
}

// 推送消息
func (ms *mySvr) Message(ctx context.Context, mreq *myserver.MessageReq, mrsp *myserver.MessageRsp) (err error) {
	logger.Debugf("in Message => mreq user id: %v", mreq.UserId)

	mrsp.Confirm = true
	mrsp.Host = PORT

	return
}

func main() {
	var ip string
	flag.StringVar(&ip, "ip", "156", "ip piece, like 156 in 10.0.0.156")
	flag.Parse()

	tcpSrv := micro.NewService(
		micro.Name("grpc-tcp"),
		micro.Address(fmt.Sprintf("10.0.0.%v:7777", ip)),
		micro.Registry(consul.NewRegistry(registry.Addrs(cs))),
		//micro.Transport(grpc.NewTransport()),
		//micro.Transport(tcp.NewTransport()),
	)

	logger.Debugf("server : %v", tcpSrv.Server().Options().Id)

	myserver.RegisterMyServerHandler(tcpSrv.Server(), new(mySvr))

	if err := tcpSrv.Run(); err != nil {
		logger.Errorf("tcpSrv execution failed: %v", err)
	}
}

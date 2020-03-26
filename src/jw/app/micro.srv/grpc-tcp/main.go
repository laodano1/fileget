package main

import (
	"app/micro.srv/grpc-tcp/proto/myserver"
	"context"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/tcp"
)

var (
	logger = golog.New("grpc-tcp")
	PORT   = ":7777"
	cs = "10.0.0.60:8500"
)

type mySvr struct{}

func (ms *mySvr) ConnectA(ctx context.Context, req *myserver.ConnectReq, rsp *myserver.ConnectRsp) (err error) {
	logger.Debugf("in ConnectA => req user id: %v", req.UserId)

	rsp.Confirm = true

	return
}
// 用户断线
func (ms *mySvr) Disconnect(ctx context.Context, dreq *myserver.DisconnectReq, drsp *myserver.DisconnectRsp) (err error) {
	logger.Debugf("in Disconnect => dreq user id: %v", dreq.UserId)

	drsp.Confirm = true

	return
}

// 推送消息
func (ms *mySvr) Message(ctx context.Context, mreq *myserver.MessageReq, mrsp *myserver.MessageRsp) (err error) {
	logger.Debugf("in Message => mreq user id: %v", mreq.UserId)

	mrsp.Confirm = true

	return
}

func main() {
	tcpSrv := micro.NewService(
		micro.Name("grpc-tcp"),
		micro.Address(PORT),
		micro.Registry(consul.NewRegistry(registry.Addrs(cs))),
		//micro.Transport(grpc.NewTransport()),
		micro.Transport(tcp.NewTransport()),
	)

	logger.Debugf("server: %v", tcpSrv.Server().Options().Id)

	myserver.RegisterMyServerHandler(tcpSrv.Server(), new(mySvr))

	if err := tcpSrv.Run(); err != nil {
		logger.Errorf("tcpSrv execution failed!")
	}
}

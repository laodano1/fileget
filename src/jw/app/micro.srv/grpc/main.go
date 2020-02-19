package main

import (
	"app/micro.srv/grpc/proto/aaa"
	"context"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro"
)

type myTServerHandler struct{}

//ConnectA(context.Context, MyServer_ConnectAStream) error
//// 用户上线b
//ConnectB(context.Context, *ConnectReq, MyServer_ConnectBStream) error
//// 用户上线c
//ConnectC(context.Context, MyServer_ConnectCStream) error
//// 用户断线
//Disconnect(context.Context, *DisconnectReq, *DisconnectRsp) error
//// 推送消息
//Message(context.Context, *MessageReq, *MessageRsp) error

func (msh *myTServerHandler) ConnectA(ctx context.Context, cas aaa.MyServer_ConnectAStream) (err error) {

	return
}

func (msh *myTServerHandler) ConnectB(ctx context.Context, creq *aaa.ConnectReq, cas aaa.MyServer_ConnectBStream) (err error) {

	return
}

func (msh *myTServerHandler) ConnectC(ctx context.Context, cas aaa.MyServer_ConnectCStream) (err error) {

	return
}

func (msh *myTServerHandler) Disconnect(ctx context.Context, dreq *aaa.DisconnectReq, drsp *aaa.DisconnectRsp) (err error) {

	return
}

func (msh *myTServerHandler) Message(ctx context.Context, mreq *aaa.MessageReq, mrsp *aaa.MessageRsp) (err error) {

	return
}

var (
	logger = golog.New("grpc stream")
	PORT   = "0.0.0.0:7777"
)

func main() {
	service := micro.NewService(
		micro.Address(PORT),
		//micro.Registry(consul.NewRegistry(registry.Addrs())),
		//micro.Registry(etcd.NewRegistry(registry.Addrs())),
		//micro.Registry(kubernetes.NewRegistry(registry.Addrs())),
	)

	aaa.RegisterMyServerHandler(service.Server(), new(myTServerHandler))

	if err := service.Run(); err != nil {
		logger.Errorf("service run failed: %v", err)
	}
}

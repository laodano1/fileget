package main

import (
	"app/micro.srv/grpc/proto/aaa"
	"context"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/tcp"
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
	PORT   = ":7777"
	cs = "10.0.0.60:8500"
)

func GetGWIP(name, chkid string) {
	//cfg := api.DefaultConfig()
	//cfg.Address = cs
	//cli, _ := api.NewClient(cfg)


	//ms, _ := cli.Agent().Services()
	//for sName, _ := range ms {
	//	logger.Debugf("1. agent check: %v", sName)
	//
	//}

	//mp, _ := cli.Agent().Checks()
	//for chkId, _ := range mp {
	//	logger.Debugf("1. agent check id: %v", chkId)
	//	//logger.Debugf("1. agent check id: %v, %v", chkId, ac)
	//	//if ac.CheckID == fmt.Sprintf("%v-%v", name, chkid) {
	//	//	logger.Debugf("2. agent check: %v", ac)
	//	//}
	//}

}

func main() {
	service := micro.NewService(
		micro.Name("grpc test"),
		micro.Address(PORT),
		micro.Registry(consul.NewRegistry(registry.Addrs(cs))),
		//micro.Transport(grpc.NewTransport()),
		micro.Transport(tcp.NewTransport()),

		//micro.Registry(etcd.NewRegistry(registry.Addrs())),
		//micro.Registry(kubernetes.NewRegistry(registry.Addrs())),
	)

	logger.Debugf("server: %v", service.Server().Options().Id)
	//GetGWIP(service.Name(), service.Server().Options().Id)

	aaa.RegisterMyServerHandler(service.Server(), new(myTServerHandler))

	if err := service.Run(); err != nil {
		logger.Errorf("service run failed: %v", err)
	}
}

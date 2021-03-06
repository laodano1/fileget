package main

import (
	message "app/cellnetrpc/proto"
	"app/cellnetrpc/server"
	_ "app/cellnetrpc/server"
	"context"
	"encoding/json"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"log"
	"net"
	"reflect"
	"sync"
)

type TT struct {
	UserId  int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	WsId    string `protobuf:"bytes,2,opt,name=WsId,proto3" json:"WsId,omitempty"`
	Service string `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
}

type websocketServer struct{}

//// 用户上线
//Connect(context.Context, *ConnectReq) (*ConnectRsp, error)
//// 用户断线
//Disconnect(context.Context, *DisconnectReq) (*DisconnectRsp, error)
//// 推送消息
//Message(context.Context, *MessageReq) (*MessageRsp, error)

func (wss *websocketServer) Connect(context.Context, *message.ConnectReq) (rsp *message.ConnectRsp, err error) {
	log.Println("in connect")

	return
}

func (wss *websocketServer) Disconnect(context.Context, *message.DisconnectReq) (drsp *message.DisconnectRsp, err error) {
	log.Println("in disconnect")
	return
}

func (wss websocketServer) Message(ctx context.Context, req *message.MessageReq) (mrsp *message.MessageRsp, err error) {
	log.Printf("token: %v, userId: %v, %v", req.Token, req.UserId, req)
	tmp := new(TT)
	err = json.Unmarshal(req.Data, tmp)
	if err != nil {
		log.Fatalf("unmarshal data failed: %v", err)
		return
	}

	log.Printf("in Message: %v", reflect.TypeOf(tmp))
	mrsp = &message.MessageRsp{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return
}

//var RpcServer *server.serverPeer

func CreatePeer(service, id, peerType, procName, host string) cellnet.GenericPeer {
	q := cellnet.NewEventQueue()
	p := peer.NewGenericPeer(service, id, host, q)
	//p := peer.NewGenericPeer("rpc.Server", "rcp", ":8801", q)

	//svr := p.(server.Server).Server()
	//message.RegisterWebsocketServer(svr, &websocketServer{})

	//q.StartLoop()
	//q.Wait()

	return p

}

func main() {
	SvrName, Id, Type, ProtoType, Host := "rpc.Server", "rpc", "", "", ":8801"
	p := CreatePeer(SvrName, Id, Type, ProtoType, Host)

	p.Start()
	svr := p.(server.Server).Server()
	lis, err := net.Listen("tcp", Host)

	if err != nil {
		return
	}

	message.RegisterWebsocketServer(svr, &websocketServer{})

	var sg sync.WaitGroup
	sg.Add(1)
	go func() {
		sg.Done()
		svr.Serve(lis)
	}()

	sg.Wait()

	log.Println("gRPC server start to listen on", Host)
	p.Queue().StartLoop()
	p.Queue().Wait()

}

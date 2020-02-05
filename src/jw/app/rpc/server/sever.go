package main

import (
	"app/rpc/proto/app/gameprovide"
	"context"
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/gogopb"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/davyxu/cellnet/rpc"
	"github.com/davyxu/cellnet/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"reflect"
	"time"
)

const peerAddress = "127.0.0.1:17701"

type TestEchoACK struct {
	Msg   string
	Value int32
}

func (self *TestEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

//func (self *TestEchoACK) ProtoMessage()  {}
//func (self *TestEchoACK) Reset()         { *self = TestEchoACK{} }

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		//Codec: codec.MustGetCodec("gogopb"),
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*TestEchoACK)(nil)).Elem(),
		ID:    int(util.StringHash("main.TestEchoACK")),
	})
}

func server() {
	queue := cellnet.NewEventQueue()

	peerIns := peer.NewGenericPeer("tcp.Acceptor", "rpc server", peerAddress, queue)

	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted: // 接受一个连接
			fmt.Println("server accepted session")

		case *rpc.RecvMsgEvent:
			log.Printf("msg: %v", msg.Session().ID())
			ev.(*rpc.RecvMsgEvent).Reply(&TestEchoACK{
				Msg:   "aabbcc",
				Value: 0,
			})

		//case *TestEchoACK: // 收到连接发送的消息
		//
		//	fmt.Printf("server recv %+v\n", msg)
		//
		//	ack := &TestEchoACK{
		//		Msg:   msg.Msg,
		//		Value: msg.Value,
		//	}
		//
		//	// 当服务器收到的是一个rpc消息
		//	if rpcevent, ok := ev.(*rpc.RecvMsgEvent); ok {
		//
		//		// 以RPC方式回应
		//		rpcevent.Reply(ack)
		//	} else {
		//
		//		// 收到的是普通消息，回普通消息
		//		ev.Session().Send(ack)
		//	}

		case *cellnet.SessionClosed: // 连接断开
			fmt.Println("session closed: ", ev.Session().ID())
		}

	})

	log.Printf("server started...")
	peerIns.Start()

	queue.StartLoop()
	queue.Wait()
}

type msgServer struct{}

func (s *msgServer) Connection(ctx context.Context, req *gameprovide.Request) (res *gameprovide.Response, err error) {

	return
}

func grpcServer2() {

	queue := cellnet.NewEventQueue()

	peerIns := peer.NewGenericPeer("tcp.Acceptor", "rpc server", peerAddress, queue)

	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {

	})

	tcpA := peerIns.(cellnet.TCPAcceptor)

	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}

	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(2 * time.Second),
	}

	gs := grpc.NewServer(opts...)
	gameprovide.RegisterCommunicationServer(gs, &msgServer{})

	reflection.Register(gs)
	if err := gs.Serve(); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}

}

func main() {
	//server()

}

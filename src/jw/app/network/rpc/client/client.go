package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/gogopb"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/davyxu/cellnet/rpc"
	"github.com/davyxu/cellnet/util"
	"log"
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

func clientHandler(event cellnet.Event) {
	switch msg := event.Message().(type) {
	case *cellnet.SessionClosed:
		log.Println("session closed")

	case *cellnet.SessionConnected:
		log.Println("session connected")

	case *TestEchoACK:
		log.Printf("client received %v", msg.String())
	}
}

func clientSyncRPC() {

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("tcp.Connector", "async rpc", peerAddress, queue)

	// 创建一个消息同步接收器
	rv := proc.NewSyncReceiver(p)

	proc.BindProcessorHandler(p, "tcp.ltv", rv.EventCallback())

	p.Start()

	queue.StartLoop()

	// 等连接上时
	rv.WaitMessage("cellnet.SessionConnected")

	// 同步RPC
	rpc.CallSync(p, &TestEchoACK{
		Msg:   "hello",
		Value: 1234,
	}, time.Second)

	//rv.Recv()
	//time.Sleep(2 * time.Second)
}

func main() {
	clientSyncRPC()
}

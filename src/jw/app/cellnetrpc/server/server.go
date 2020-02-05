package server

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type serverPeer struct {
	svr *grpc.Server
	*peer.CorePeerProperty
	options []grpc.ServerOption
	lis     net.Listener
	sg      sync.WaitGroup
}

// 开启端，传入地址
func (self *serverPeer) Start() cellnet.Peer {
	self.svr = grpc.NewServer(self.options...)
	//lis, err := net.Listen("tcp", self.Address())
	//
	//if err != nil {
	//	return nil
	//}

	//self.lis = lis
	//self.sg.Add(1)
	//go func() {
	//	self.sg.Done()
	//	self.svr.Serve(lis)
	//}()
	//
	//self.sg.Wait()

	return self
}

// 停止通讯端
func (self *serverPeer) Stop() {
	if self.svr != nil {
		self.svr.Stop()
		self.svr = nil
	}
}

// Peer的类型(protocol.type)，例如tcp.Connector/udp.Acceptor
func (self *serverPeer) TypeName() string {
	return "rpc.Server"
}

//grpc server
func (self *serverPeer) Server() *grpc.Server {
	return self.svr
}

//选项
func (self *serverPeer) Options() []grpc.ServerOption {
	return self.options
}

//设置选项
func (self *serverPeer) SetOption(option ...grpc.ServerOption) {
	self.options = option
}

func newRcpServer() cellnet.Peer {
	p := new(serverPeer)
	p.CorePeerProperty = &peer.CorePeerProperty{}
	return p
}

func init() {
	peer.RegisterPeerCreator(newRcpServer)
}

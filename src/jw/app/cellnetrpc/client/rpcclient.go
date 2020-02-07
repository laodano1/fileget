package main

import (
	message "app/cellnetrpc/proto"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"time"
)

type TT struct {
	UserId  int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	WsId    string `protobuf:"bytes,2,opt,name=WsId,proto3" json:"WsId,omitempty"`
	Service string `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
}

func main() {
	Host := ":8801"
	conn, err := grpc.Dial(Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := message.NewWebsocketClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data := &TT{
		UserId:  11,
		WsId:    "123",
		Service: "rpc-123",
		Token:   "qwe1234455rrt==",
	}
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("cannot marshal data: %v", err)
	}

	mreq := &message.MessageReq{
		UserId: 999,
		Data:   d,
		Token:  "xxx22323frger",
	}

	log.Printf("send to server: %v", mreq)
	_, err = c.Message(ctx, mreq)
	if err != nil {
		log.Fatalf("cannot send message: %v", err)
	}

	log.Println("response from server")

	//_, err = c.Disconnect(ctx, &message.DisconnectReq{UserId: 222, Token: "32reraef23r"})

}

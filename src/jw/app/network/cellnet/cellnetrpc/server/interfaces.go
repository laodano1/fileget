package server

import "google.golang.org/grpc"

type Server interface {
	//grpc server
	Server() *grpc.Server

	//选项
	Options() []grpc.ServerOption

	//设置选项
	SetOption(option ...grpc.ServerOption)
}

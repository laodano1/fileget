package main

import "google.golang.org/grpc"

type dockPlatformgPRC struct {
	grpcSrv *grpc.Server
}

func NewdockPlatformgPRC() *dockPlatformgPRC {
	return &dockPlatformgPRC{}
}

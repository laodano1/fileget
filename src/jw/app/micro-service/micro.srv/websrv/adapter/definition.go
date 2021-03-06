package main

type newArchAdapter struct {
	UserWsSrv        *towardUserWsSrv
	dockPlatformgPRC *dockPlatformgPRC
	recvCh           *msgTransit
	sendCh           *msgTransit
}

const (
	wsCloseAbnormal = iota
	wsCloseNormal
)

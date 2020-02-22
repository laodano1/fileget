package main

type newArchAdapter struct {
	UserWsSrv        *towardUserWsSrv
	BackendWsCli     *towardBackendWsCli
	dockPlatformgPRC *dockPlatformgPRC
	recvCh           *msgTransit
	sendCh           *msgTransit
}

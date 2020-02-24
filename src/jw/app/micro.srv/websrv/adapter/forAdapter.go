package main

func NewAdapter() *newArchAdapter {
	return &newArchAdapter{
		UserWsSrv:        NewWSSrv(),
		dockPlatformgPRC: NewdockPlatformgPRC(),
		recvCh:           NewMsgTransit(),
		sendCh:           NewMsgTransit(),
	}
}

func (naa *newArchAdapter) Init() {
	//lg.Debugf("connecting to %s", u.String())
	naa.UserWsSrv.Init()
	naa.UserWsSrv.Start()

}

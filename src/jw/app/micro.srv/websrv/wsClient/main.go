package main

import (
	"github.com/davyxu/golog"
)

var (
	srvAdd = "10.0.0.156"
)

var (
	lg   = golog.New("ws client")
	PORT = "0.0.0.0:80"
)

func NewAdapter() *newArchAdapter {
	return &newArchAdapter{
		UserWsSrv:        NewWSSrv(),
		BackendWsCli:     NewWSCli(),
		dockPlatformgPRC: NewdockPlatformgPRC(),
		recvCh:           NewMsgTransit(),
		sendCh:           NewMsgTransit(),
	}
}

func (naa *newArchAdapter) Init() {
	//lg.Debugf("connecting to %s", u.String())

}

func main() {
	lg.SetParts()

	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

}

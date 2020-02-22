package main

import (
	"github.com/gorilla/websocket"
	"net/url"
)

type towardBackendWsCli struct {
	dialer *websocket.Dialer
}

func NewWSCli() *towardBackendWsCli {
	return &towardBackendWsCli{
		dialer: websocket.DefaultDialer,
	}
}

func (wsCli *towardBackendWsCli) ReadAndWriteWithBackend() {
	u := url.URL{
		Scheme: "ws",
		Host:   srvAdd,
		Path:   "/game",
	}
	c, _, err := wsCli.dialer.Dial(u.String(), nil)
	if err != nil {
		lg.Debugf("dial:", err)
	}

	c.SetCloseHandler(func(code int, text string) (err error) {

		return
	})
	//for {
	//	//c.ReadMessage()
	//
	//	//c.WriteMessage()
	//
	//}
}

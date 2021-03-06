package main

import (
	"github.com/davyxu/golog"
	"github.com/gorilla/websocket"
	"net/url"
)

func main() {
	logger := golog.New("ws-client")

	host := "web2-sit.wonderland.life"
	u := url.URL{Scheme: "ws", Host: host, Path: "/2gLv1z7it6cm8kwA/b3a710101377748895851377748897a5ae274d55c5d4f3faa983da4999ea9039/600101"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		logger.Errorf("ws dial failed: %v", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Errorf("ws read failed:", err)
			return
		}

		logger.Debugf("ws received: %s\n", message)
	}
}

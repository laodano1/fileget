package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

type towardBackendWsCli struct {
	conn *websocket.Conn
}

func NewWSCli(host, path string) (cli *towardBackendWsCli, err error) {
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   path, //"/game",
	}
	var c *websocket.Conn
	var resp *http.Response
	c, resp, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		lg.Errorf("dial(%v) error:", u.String(), err)
		return
	}

	lg.Debugf("NewWSCli | http response: %v", resp)

	c.SetCloseHandler(func(code int, text string) (err error) {
		switch code {
		case websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
			lg.Errorf("websocket closed with error: %v", text)

		case websocket.CloseNormalClosure:
			lg.Infof("websocket close successfully: %v", text)

		default:
			lg.Infof("in SetCloseHandler default branch!")
		}
		return
	})
	cli = &towardBackendWsCli{
		conn: c,
	}
	return
}

func (wsCli *towardBackendWsCli) readAFromBackend() (mt int, data []byte, err error) {
	mt, data, err = wsCli.conn.ReadMessage()
	return
}

func (wsCli *towardBackendWsCli) write2Backend(mt int, data []byte) (err error) {
	err = wsCli.conn.WriteMessage(mt, data)
	return
}

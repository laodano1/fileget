package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
	"time"
)

type towardUserWsSrv struct {
	srv web.Service
}

func homepage(c *gin.Context) {
	c.String(http.StatusOK, "New Adapter works well!")
}

func gameOperation(c *gin.Context) {
	wsUpgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) (ok bool) {
			return true
		},
	}

	wsConn, err := wsUpgrader.Upgrade(c.Writer, c.Request, http.Header{})
	if err != nil {
		lg.Errorf("upgrade failed: %v", err)
	}
	//defer wsConn.Close()

	done := make(chan int)

	wsConn.SetCloseHandler(func(code int, text string) (err error) {
		switch code {
		case websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
			lg.Errorf("websocket closed with error: %v", text)
			done <- wsCloseAbnormal
		case websocket.CloseNormalClosure:
			lg.Infof("websocket close successfully: %v", text)
			done <- wsCloseNormal
		default:
			lg.Infof("in SetCloseHandler default branch!")
		}
		return
	})

	wsConn.SetPingHandler(func(appData string) (err error) {
		if err := wsConn.WriteMessage(websocket.TextMessage, []byte(appData)); err != nil {
			lg.Errorf("PingHandler write message error: %v", err)
		}
		lg.Infof("in SetPingHandler")
		return
	})

	wsConn.SetPongHandler(func(appData string) (err error) {
		if err := wsConn.WriteMessage(websocket.TextMessage, []byte(appData)); err != nil {
			lg.Errorf("PongHandler write message error: %v", err)
		}
		lg.Infof("in SetPongHandler")
		return
	})

	cli, err := NewWSCli("10.0.0.146", "/game")
	if err != nil {
		lg.Errorf("new websocket dial failed: %v", err)
	}

	readMsgOperation(wsConn, cli)
	writeMsgOperation(wsConn, cli)
}

func readMsgOperation(wsConn *websocket.Conn, cli *towardBackendWsCli) {
	for {
		mt, message, err := wsConn.ReadMessage()
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
				lg.Errorf("ws close with error: %v", ce.Text)
			case websocket.CloseMessage:
				lg.Infof("socket closed!")
			}
			return
		}

		lg.Infof("client sent message: '%s', mt: %v", fmt.Sprintf("server received data: %v", string(message)), mt)

		cli.readAFromBackend()
		// TODO: write to recv channel

	}
}

func writeMsgOperation(wsConn *websocket.Conn, cli *towardBackendWsCli) {
	for {
		msg := []byte{}
		err := wsConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			lg.Errorf("WriteMessage failed: %v", err)
		}

		cli.write2Backend()
		// TODO: write to send channel

	}
}

func NewWSSrv() *towardUserWsSrv {
	s := web.NewService(
		web.Name("game-600101"),
		web.Address(PORT),
		web.Registry(
			consul.NewRegistry(
				registry.Addrs("10.0.0.146"),
			),
		),
		web.RegisterTTL(10*time.Second),
	)
	return &towardUserWsSrv{
		srv: s,
	}
}

func (wsSrv *towardUserWsSrv) Init() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	router.GET("/", homepage)
	router.GET("/game", gameOperation)

	wsSrv.srv.Handle("/", router)
}

func (wsSrv *towardUserWsSrv) Start() {
	go func() {
		defer func() {
			recover()
		}()
		if err := wsSrv.srv.Run(); err != nil {
			lg.Errorf("newArchAdapter gin server run error: ", err)
		}
	}()

	//naa.ReadAndWriteWithBackend()
}

package main

import (
	"fmt"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"google.golang.org/grpc"
	"net/http"
	"net/url"
	"time"
)

var (
	srvAdd = "10.0.0.156"
)

type newArchAdapter struct {
	TowardPlayer  web.Service
	TowardBackEnd *websocket.Dialer
	ForPlatform   *grpc.Server
}

var (
	lg   = golog.New("ws client")
	PORT = "0.0.0.0:80"
)

func NewAdapter() *newArchAdapter {

	//lg.Debugf("connecting to %s", u.String())
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

	return &newArchAdapter{
		TowardPlayer:  s,
		TowardBackEnd: websocket.DefaultDialer,
	}
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

	wsConn.SetCloseHandler(func(code int, text string) (err error) {
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

		lg.Infof("client sent message: '%s'", fmt.Sprintf("server received data: %v", string(message)))

		//err = wsConn.WriteMessage(mt, []byte(fmt.Sprintf("server received data: %v", message)))
		err = wsConn.WriteMessage(mt, message)
		if err != nil {
			lg.Errorf("WriteMessage failed: %v", err)
		}
	}
}

func (naa *newArchAdapter) Init() {
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

	naa.TowardPlayer.Handle("/", router)

}

func (naa *newArchAdapter) Start() {
	go func() {
		defer func() {
			recover()
		}()
		if err := naa.TowardPlayer.Run(); err != nil {
			lg.Errorf("newArchAdapter gin server run error: ", err)
		}
	}()

	naa.ReadAndWriteWithBackend()
}

func (naa *newArchAdapter) ReadAndWriteWithBackend() {
	u := url.URL{
		Scheme: "ws",
		Host:   srvAdd,
		Path:   "/game",
	}
	c, _, err := naa.TowardBackEnd.Dial(u.String(), nil)
	if err != nil {
		lg.Debugf("dial:", err)
	}

	c.SetCloseHandler(func(code int, text string) (err error) {

		return
	})
	for {
		c.ReadMessage()

		c.WriteMessage()

	}
}

func main() {
	lg.SetParts()

	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

}

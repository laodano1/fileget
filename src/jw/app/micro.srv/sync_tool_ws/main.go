package main

import (
	"fileget/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"strings"
)

var (
	PORT = ":9999"
)

func syncToolPage(c *gin.Context) {
	wsUpgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) (ok bool) {
			return true
		},
	}

	wsConn, err := wsUpgrader.Upgrade(c.Writer, c.Request, http.Header{})
	if err != nil {
		util.Lg.Errorf("upgrade failed: %v", err)
	}

	wsConn.SetCloseHandler(func(code int, text string) (err error) {
		switch code {
		case websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
			util.Lg.Errorf("websocket closed with error: %v", text)

		case websocket.CloseNormalClosure:
			util.Lg.Infof("websocket close successfully: %v", text)
		default:
			util.Lg.Infof("in SetCloseHandler default branch!")
		}
		return
	})

	wsConn.SetPingHandler(func(appData string) (err error) {
		if err := wsConn.WriteMessage(websocket.TextMessage, []byte(appData)); err != nil {
			util.Lg.Errorf("PingHandler write message error: %v", err)
		}
		util.Lg.Infof("in SetPingHandler")
		return
	})

	wsConn.SetPongHandler(func(appData string) (err error) {
		if err := wsConn.WriteMessage(websocket.TextMessage, []byte(appData)); err != nil {
			util.Lg.Errorf("PongHandler write message error: %v", err)
		}
		util.Lg.Infof("in SetPongHandler")
		return
	})

	for {
		mt, message, err := wsConn.ReadMessage()
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr:
				util.Lg.Errorf("ws close with error: %v", ce.Text)
			case websocket.CloseMessage:
				util.Lg.Infof("socket closed!")
			}
			return
		}

		util.Lg.Infof("client sent message: '%s'", fmt.Sprintf("server received data: %v", string(message)))

		//err = wsConn.WriteMessage(mt, []byte(fmt.Sprintf("server received data: %v", message)))
		err = wsConn.WriteMessage(mt, message)
		if err != nil {
			util.Lg.Errorf("WriteMessage failed: %v", err)
		}
	}

}

func homepage(c *gin.Context) {
	//c.String(http.StatusOK, "hello world home page")
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"SplitFileName": func(fullStr string) (subStr string) {
			return strings.Split(fullStr, ".")[0]
		},
	}).Parse(tmplStr)) // Create a template

	htmlRender := render.HTML{
		Template: tmpl,
		Name:     "",
		Data:     nil,
	}

	c.Render(http.StatusOK, htmlRender)
}

func main() {
	e := gin.Default()

	e.GET("/sync", syncToolPage)
	e.GET("/", homepage)

	if err := e.Run(PORT); err != nil {
		util.Lg.Errorf("web server run failed: %v", err)
	}
}

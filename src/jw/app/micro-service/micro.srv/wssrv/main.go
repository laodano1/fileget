package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
	"html/template"
	"net/http"
	"strings"
)

var (
	SERVERNAME = "go.micro.websocket"
	VERSION    = "websocket v0.0.1"
	//PORT       = "0.0.0.0:9999"
	PORT = ":9999"
	lg   = golog.New(SERVERNAME)
)

var tmplStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

</head>
<body>
	<div>
		hello world in template
	</div>
	<input id="input" type="text" />
	<button onclick="send()">Send</button>
	<button onclick="connectttt()">Connect</button>
	<button onclick="clooose()">Close</button>
	<pre id="output"></pre>
	<script>
		var input  = document.getElementById("input");
		var output = document.getElementById("output");
		var socket = new WebSocket("ws://localhost:9999/ws");
		startt()

		function startt() {
			socket.onopen = function (e) {
				output.innerHTML = "Status: Connected\n";
			};
		
			socket.onmessage = function (e) {
				output.innerHTML += "Server: " + e.data + "\n";
			};
	
			socket.onclose = function (e) {
				output.innerHTML = "session closed!";
				console.log("onclose invoked!")
			};
	
			socket.onerror = function(event) {
				console.log("onerror invoked!" + e)
			};
		}
		

		function send() {
			socket.send(input.value);
			console.log("send() invoked!")
			input.value = "";
		};

		function clooose() {
			socket.close();
			output.value = "session closed!";
		};

		function connectttt() {
			console.log("connectttt() invoked!")
			switch (socket.readyState) {
			  case WebSocket.CONNECTING:
				// do something
				break;
			  case WebSocket.OPEN:
				// do something
				break;
			  case WebSocket.CLOSING:
				// do something
				break;
			  case WebSocket.CLOSED:
				// do something
				socket = new WebSocket("ws://localhost:9999/ws");
				startt()
				break;
			  default:
				// this never happens
				break;
			}
			
			output.value = "session connected!";
		};

	</script>
</body>
`

func wsFunc(c *gin.Context) {
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
	ws := web.NewService(
		web.Name(SERVERNAME),
		web.Address(PORT),
		web.Version(VERSION),
	)

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
	router.GET("/ws", wsFunc)

	ws.Handle("/", router)

	if err := ws.Run(); err != nil {
		lg.Errorf("web server run failed: %v", err)
	}

}

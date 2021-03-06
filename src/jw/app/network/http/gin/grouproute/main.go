package main

import (
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

type myWebServer struct {
	name string
	wsrv web.Service
}

var lg = golog.New("my-web-server")

func NewWebServer() *myWebServer {
	mws := web.NewService(
		web.Name("my-web-service"),
	)
	return &myWebServer{
		name: "haha",
		wsrv:  mws,
	}
}

func homepage(c *gin.Context)  {
	c.String(http.StatusOK, "hello world, this is homepage!")
}

func main() {
	if err := NewWebServer().wsrv.Run(); err != nil {
		lg.Errorf("run failed!", err)
	}
}

func (mws *myWebServer) InitWSrv() (err error) {
	gin.ForceConsoleColor()
	e := gin.Default()
	e.Use(gin.Recovery())

	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	e.GET("/", homepage)

	mws.wsrv.Handle("/", e)

	if err = mws.wsrv.Run(); err != nil {
		lg.Errorf("web server run failed: %v", err)
		return
	}

	return
}

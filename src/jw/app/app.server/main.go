package main

import (
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
)

const (
	srvAdd = ":9999"
)

func haha(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world in gin!")
}

func homepage(ctx *gin.Context) {
	tmpl := template.Must(template.New("aaa").Parse(hptemplate))

	dt := gin.H{
		"title": "hello world",
		"log": "this is log",
	}

	rd := render.HTML{
		Template: tmpl,
		Name:     "aaa",
		Data:     dt,
	}
	ctx.Render(http.StatusOK, rd)
	//ctx.String(http.StatusOK, "this is homepage!")
}

func main() {
	logger := golog.New("my-web-server")
	//logger.SetParts()

	router := gin.Default()
	router.GET("/", homepage)
	router.GET("/haha", haha)

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	if err := router.Run(srvAdd); err != nil {
		logger.Errorf("router run failed: %v", err)
	}
	//
	//
	//go util.ShowMemStat(10, logger)
	//
	//// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	//queue := cellnet.NewEventQueue()
	//
	////监听端口
	//p := peer.NewGenericPeer("gorillaws.Acceptor", "my-app-server", srvAdd, queue)
	//
	////绑定事件处理
	//// 绑定固定回调处理器, procName-"第二个参数"来源于RegisterProcessor注册的处理器，形如: 'gorillaws.ltv'
	//proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
	//	switch msg := ev.Message().(type) {
	//	////连接上了
	//	//case *cellnet.SessionConnected:
	//	//	logger.Infof("Session(%v) connected.", ev.Session().ID())
	//
	//	//接受客户端连接过来
	//	case *cellnet.SessionAccepted:
	//		logger.Infof("Session(%v) Accepted, msg: %v", ev.Session().ID(), msg)
	//
	//	//会话断开
	//	case *cellnet.SessionClosed:
	//		logger.Infof("Session(%v) Closed", ev.Session().ID())
	//
	//		//接受别的消息类型
	//		//case :
	//
	//	}
	//
	//})
	//
	////开始监听
	//p.Start()
	//
	////开启事件队列循环
	//queue.StartLoop()
	//
	//// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	//queue.Wait()

}

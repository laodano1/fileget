package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"
	"jw/common/util"
	"runtime"

	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

var (
	srvAdd = ":8888"
	logger *golog.Logger
	memStat *runtime.MemStats
)


func main()  {
	logger = golog.New("my.robot")
	logger.SetParts()
	go util.ShowMemStat(10, logger)

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Connector", "my-robot-cli", srvAdd, queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			logger.Infof("session(%v) connected, %v", ev.Session().ID(), msg.SystemMessage)

		case *cellnet.SessionClosed:
			logger.Infof("Session(%v) closed", ev.Session().ID())
		}
	})

	p.Start()

	queue.StartLoop()

	queue.Wait()

}




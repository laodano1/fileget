package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"
	"runtime"
)

var (
	srvAdd = ":8888"
	logger *golog.Logger
	memStat *runtime.MemStats
)


func main()  {
	logger = golog.New("my.robot")
	logger.SetParts()

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Connector", "my-robot-cli", srvAdd, queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:

		case *cellnet.SessionClosed:

		}
	})

	p.Start()

	queue.StartLoop()

	queue.Wait()

}




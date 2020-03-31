package main

import (
	"github.com/davyxu/golog"
	"runtime"
	"time"
)

var logger = golog.New("basic-pool")

func main() {
	p := newPool()

	p.initPool(10)
	p.takeWork(100)
	p.isFullRollCall()
	time.AfterFunc(10 * time.Second, func() {
		for i := 0; i < 10; i++ {
			p.OffDuty <- true
		}
	})

	time.Sleep(300 * time.Microsecond)
	logger.Debugf("goroutine number: '%v'", runtime.NumGoroutine())

	<- p.Done
	logger.Debugf("all workers leave!")
}

package main

import (
	"github.com/davyxu/golog"
	"sync"
	"time"
)

var flag = 0
var logger = golog.New("synce-once")

func start(id int, oc *sync.Once, wg *sync.WaitGroup) {
	wg.Add(1)
	time.Sleep(100 * time.Microsecond)
	oc.Do(func() {
		logger.Debugf("'%v' set flag", id)
		flag = 1
	})
	logger.Debugf("in goroutine: %v", id)
	wg.Done()
}

func main() {
	logger.EnableColor(true)
	//var once sync.Once
	//var wg sync.WaitGroup

	pl := sync.Pool{}

	logger.Debugf("pool item: %v", pl.Get())

	for i := 1; i <= 5; i++ {
		//go start(i, &once, &wg)
		logger.Debugf("pool put item: %v", i)
		pl.Put(i)
	}
	//wg.Wait()
	logger.Debugf("================================")
	for i := 5; i > 0; i-- {
		//go start(i, &once, &wg)
		//pl.Put(i)
		logger.Debugf("pool get item: %v", pl.Get())
	}

	var t2 sync.Cond

}

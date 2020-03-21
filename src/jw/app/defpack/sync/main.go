package main

import (
	"github.com/davyxu/golog"
	"sync"
	"time"
)


var logger = golog.New("sync-cond")

func worker(c *sync.Cond, id int, wg *sync.WaitGroup) {
	wg.Add(1)
	c.L.Lock()
	logger.Debugf("'%v' wait", id)
	c.Wait()
	logger.Debugf("'%v' start", id)
	c.L.Unlock()
	wg.Done()

}


func s1(c *sync.Cond, id int, wg *sync.WaitGroup)  {
	wg.Add(1)
	logger.Debugf("'s-%v' before sleep", id)
	time.Sleep(1 * time.Second)
	c.Signal()
	//c.Broadcast()
	logger.Debugf("'s-%v' after sleep, broadcast", id)
	wg.Done()
}

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go worker(cond, i, wg)
	}

	for i := 0; i < 5; i++ {
		go s1(cond, 9, wg)
		time.Sleep(1 * time.Second)
	}

	wg.Wait()

}


package main

import (
	"fileget/util"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(6)
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			go func(idx int) {
				util.Lg.Debugf("wait(%v)", idx)
				//time.Sleep(1 * time.Second)
				<- ch
				util.Lg.Debugf("%v done", idx)
				wg.Done()
			}(i)
			time.Sleep(500 * time.Millisecond)
		}
		wg.Done()
	}()

	m := func() {
		time.Sleep(1 * time.Second)
		util.Lg.Debugf("main wait")
		close(ch)
		util.Lg.Debugf("main done")
	}

	m()
	wg.Wait()
}

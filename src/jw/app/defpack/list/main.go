package main

import (
	"github.com/davyxu/golog"
	"runtime"
	"sync"
	"time"
)

type ch struct{
	name string
	c   chan bool
}


func main() {
	lg := golog.New("mylist")

	//r := ring.New(1)
	//lg.Debugf("1 list len: %v", r.Len())
	//for _, v := range []int{1, 2, 3, 4, 5} {
	//	item := &ring.Ring{
	//		Value: func() int {
	//			return v
	//		},
	//	}
	//	r.Link(item)
	//}
	//
	//lg.Debugf("2 list len: %v", r.Len())
	lg.Debugf("now: %v", time.Now().Format(time.RFC3339Nano))

	tch := new(ch)
	tch.c = make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		lg.Debugf("in goroutine 1")
		tch.c <- true
		lg.Debugf("in goroutine 1-1")
		wg.Done()
	}()

	//time.Sleep(2 * time.Second)
	lg.Debugf("stand by")
	go func() {
		lg.Debugf("in goroutine 2")
		<- tch.c
		lg.Debugf("in goroutine 2-1")
		wg.Done()
	}()

	//time.Sleep(4 * time.Second)
	lg.Debugf("goroutine num: %v", runtime.NumGoroutine())
	wg.Wait()
	lg.Debugf("bye bye")
}

package main

import (
	"github.com/davyxu/golog"
	"sync"
)

type user struct {
	name      string
	chlock    chan bool
}

type user1 struct {
	name      string
	mlock sync.Mutex
}

func main() {
	lg := golog.New("groutine")
	lg.SetParts(golog.LogPart_Level, golog.LogPart_TimeMS, golog.LogPart_ShortFileName )
	//runtime.GOMAXPROCS(4)
	//runtime.Goexit()
	//time.After()
	//
	//var sum uint32
	//var done chan bool
	//num := 100
	//
	//for i := 0; i < num; i++ {
	//	go func() {
	//
	//	}()
	//}

	a := 1

	lg.Debugf("a: %v", a)





}

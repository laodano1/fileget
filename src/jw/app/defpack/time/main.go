package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"math/rand"
	"sync"
	"time"
)

var lg = golog.New("my-time")

func main() {

	agentId := 93
	gameId  := "600101"
	//rand.Seed(time.Now().UnixNano())

	matchId := fmt.Sprintf("%v%v%v%v", time.Now().Nanosecond(), fmt.Sprintf("%v", agentId), gameId, fmt.Sprintf("%04d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000)))
	lg.Debugf("match id: '%v'", matchId)
	//lg.Debugf("match id: '%v-%v-%v-%v'", time.Now().Nanosecond(), fmt.Sprintf("%v", agentId), gameId, fmt.Sprintf("%04d", 10000))

	//tf := TstClosure()
	tf := TstClosureWithLock()

	for i := 0; i < 100; i++ {
		go tf(i)
	}

	time.Sleep(3 * time.Second)
}

func TstClosure() func(i int) {
	var i int = 0
	return func(id int) {
		i++
		lg.Debugf("id: %v, i: %v", id, i)
	}
}

func TstClosureWithLock() func(i int) {
	var i int = 0
	var lk sync.RWMutex

	return func(id int) {
		lk.Lock()
		i++
		lg.Debugf("id: %v, i: %v", id, i)
		lk.Unlock()
	}
}
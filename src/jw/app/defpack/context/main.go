package main

import (
	"context"
	"fileget/util"
	"sync"
	"time"
)

var (
	lg = util.Lg
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10 * time.Second))
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer func () {
			wg.Done()
		}
		for {
			select {
			case <- ctx.Done():
				return
			case <- time.After(3 * time.Second):
				return
			}
		}
	}()
	wg.Wait()
	util.Lg.Debugf("bye bye!")
}

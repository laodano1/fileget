package main

import (
	"github.com/davyxu/golog"
	"sync"
)

var (
	urlPrefix = "http://whc.unesco.org"
	lg = golog.New("world-heritage-list")
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
)

func main() {
	wohelist := make(chan msg, 100)
	done     := make(chan bool)
	wg       := &sync.WaitGroup{}

	wg.Add(2)
	go getHeritageListByCountryDimension(wohelist)

	// 2 workers
	for i, _ := range []int{1, 2} {
		go func(i int) {
			for {
				select {
				case tmpMsg := <- wohelist:
					if tmpMsg.Status {
						done <- true
						wg.Done()
						return
					} else {
						//lg.Debugf("worker(%v) starts, %v", i, tmpMsg)
						GetHeritageInfo(tmpMsg.Url)
					}
				case <- done:
					wg.Done()
					return
				}
			}
		}(i)
	}

	//time.Sleep(10 * time.Second)

	wg.Wait()
	lg.Debugf("bye bye")
}

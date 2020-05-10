package main

import (
	"fileget/src/jw/app/colly/worldHL/utils"
	"fmt"
	"github.com/davyxu/golog"
	"os"
	"sync"
)

var (
	urlPrefix = "http://whc.unesco.org"
	lg = golog.New("world-heritage-list")
	exeDirPath string
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
)

func main() {
	exeDirPath, _ = utils.GetFullPathDir()
	lg.Debugf("--------- exe path: %v", exeDirPath)

	lg.SetParts(golog.LogPart_Level, golog.LogPart_Name, golog.LogPart_TimeMS)
	wohelist := make(chan msg, 1)
	done     := make(chan bool)
	wg       := &sync.WaitGroup{}

	wg.Add(3)
	go getHeritageListByCountryDimension(wohelist)

	// 2 workers
	for i, _ := range []int{1, 2, 3} {
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
						GetHeritageInfo(i,tmpMsg)
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
	utils.Write2JsonFile(allHeritageDetailList, fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, "All_World_Heritage_Detail_List.json") )
	lg.Debugf("bye bye")
}

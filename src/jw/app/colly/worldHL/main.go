package main

import (
	"fileget/util"
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"os"
	"sync"
	"time"
)

var (
	urlPrefix = "http://whc.unesco.org"
	lg = golog.New("world-heritage-list")
	exeDirPath string
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"

	allJson map[string]bool
)

func main() {
	exeDirPath, _ = util.GetFullPathDir()
	lg.Debugf("--------- exe dir path: %v", exeDirPath)

	lg.SetParts(golog.LogPart_Level, golog.LogPart_Name, golog.LogPart_TimeMS)
	lg.EnableColor(true)

	var tp string
	//flag.StringVar(&tp, "t", "country", "get type. currently support: country | whl (world heritage list) ")
	flag.StringVar(&tp, "t", "whl", "get type. currently support: country | whl (world heritage list) ")
	flag.Parse()

	wg := new(sync.WaitGroup)
	startT := time.Now()
	switch tp {
	case "country":
		getCountries()
	case "whl":
		allJson = util.GetMatchedFiles(fmt.Sprintf("%v%ctmp", exeDirPath, os.PathSeparator), "json")
		getHeritageListByCountryDimension()

		parsedCountries := make(map[string]bool)

		lk := new(sync.Mutex)
		for i, _ := range []int{1, 2, 3, 4 , 5, 6} {
			wg.Add(1)
			go parseAllHeritage(i, parsedCountries, wg, lk)
		}
	default:
	}

	wg.Wait()
	//utils.Write2JsonFile(allHeritageDetailList, fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, "All_World_Heritage_Detail_List.json") )
	lg.Debugf("total spent: %v seconds. bye bye", time.Now().Sub(startT))
}

package main

import (
	"fileget/src/jw/app/colly/worldHL/utils"
	"fmt"
	"github.com/davyxu/golog"
	"os"
	"strings"
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
	exeDirPath, _ = utils.GetFullPathDir()
	lg.Debugf("--------- exe dir path: %v", exeDirPath)

	lg.SetParts(golog.LogPart_Level, golog.LogPart_Name, golog.LogPart_TimeMS)
	lg.EnableColor(true)
	wg       := &sync.WaitGroup{}

	allJson = utils.GetMatchedFiles(fmt.Sprintf("%v%ctmp", exeDirPath, os.PathSeparator), "json")
	getHeritageListByCountryDimension()

	var parsedCountries sync.Map
	startT := time.Now()

	wg.Add(1)
	// 2 workers
	for i, _ := range []int{1} {
		go func(i int) {
			for _, c := range worldHeritages.countryOrder {
				if _, ok := parsedCountries.Load(c); !ok {  // not exist
					parsedCountries.Store(c, true)
					lg.Debugf("country: %v", c)
					for _, hl := range worldHeritages.CountryList[c].HeritageList {
						for _, oItem := range hl.Types {
							for _, h := range oItem {
								msg := parseMsg{
									Url: h.Href,
									Name: h.Name,
								}
								if !strings.Contains(msg.Url, "whc.unesco.org#") {
									GetHeritageInfo(i, msg)
									//lg.Debugf("    ---- parse heritage: %v, url: %v", msg.Name, msg.Url)
								} else {
									//lg.Debugf("    do not parse heritage: %v, url: %v", msg.Name, msg.Url)
								}
							}
						}
					}
				} else {
					continue
				}
			}
			lg.Debugf("for loop done")
			wg.Done()
		}(i)
	}

	wg.Wait()
	//utils.Write2JsonFile(allHeritageDetailList, fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, "All_World_Heritage_Detail_List.json") )
	lg.Debugf("total spent: %v seconds. bye bye", time.Now().Sub(startT))
}

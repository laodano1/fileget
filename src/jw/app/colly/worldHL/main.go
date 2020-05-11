package main

import (
	"fileget/src/jw/app/colly/worldHL/utils"
	"github.com/davyxu/golog"
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
	lg.Debugf("--------- exe dir path: %v", exeDirPath)

	lg.SetParts(golog.LogPart_Level, golog.LogPart_Name, golog.LogPart_TimeMS)
	wg       := &sync.WaitGroup{}

	getHeritageListByCountryDimension()

	var parsedCountries sync.Map

	wg.Add(2)
	// 2 workers
	for i, _ := range []int{1, 2} {
		go func(i int) {
			for _, c := range worldHeritages.countryOrder {
				if _, ok := parsedCountries.Load(c); !ok {  // not exist
					parsedCountries.Store(c, true)
					lg.Debugf("country: %v", c)
					for _, cItem := range worldHeritages.CountryList {
						for _, hItem := range cItem.HeritageList {
							for _, oItem := range hItem.Types {
								for _, h := range oItem {
									msg := parseMsg{
										Url: h.Href,
										Name: h.Name,
									}
									lg.Debugf("    parse heritage: %v, url: %v", msg.Name, msg.Url)
									GetHeritageInfo(i, msg)
								}
							}
						}
					}
				} else {
					continue
				}
			}
		}(i)
	}

	wg.Wait()
	//utils.Write2JsonFile(allHeritageDetailList, fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, "All_World_Heritage_Detail_List.json") )
	lg.Debugf("bye bye")
}

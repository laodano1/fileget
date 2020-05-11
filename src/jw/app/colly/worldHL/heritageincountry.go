package main

import (
	"fileget/src/jw/app/colly/worldHL/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
)

var (
	worldHeritages WorldHeritageList
)

func getHeritageListByCountryDimension() {
	c := colly.NewCollector(
		colly.CacheDir("./whl"),
	)
	c.UserAgent = UserAgent
	lg.Debugf("user agent: %v", c.UserAgent)
	c.OnRequest(func(req *colly.Request) {
		//lg.Debugf("request: %v", req.URL)
	})

	c.OnResponse(func(rsp *colly.Response) {
		//lg.Debugf("response...")
	})


	worldHeritages.CountryList = make(map[string]*CountryItem)

	c.OnHTML("#acc", func(e *colly.HTMLElement) {  // div
		e.DOM.ChildrenFiltered("h4").Each(func(i int, s *goquery.Selection) {  // h4
			ci := new(CountryItem)
			href, _ := s.Children().Attr("href")
			ci.Href = urlPrefix + href // e.ChildAttr("a", "href")
			ci.Name = s.Children().Text()
			ci.Type, _ = s.Attr("id")  //e.Attr("id")

			worldHeritages.CountryList[ci.Name] = ci
			worldHeritages.countryOrder = append(worldHeritages.countryOrder, ci.Name)
			//lg.Debugf("CountryItem: %v", ci)
		})

		e.DOM.ChildrenFiltered("div.list_site").Each(func(i int, s *goquery.Selection) { // div
			ht  := HeritageItem{}
			ht.Country = worldHeritages.CountryList[worldHeritages.countryOrder[i]].Name
			ht.Types     = make(map[string][]OneHeritage)
			ht.TypeOrder = make([]string, 0)
			s.ChildrenFiltered("ul").ChildrenFiltered("li").Each(func(i1 int, s1 *goquery.Selection) { // li
				//lg.Debugf("li: %v", i1)
				htp, _ := s1.Attr("class")
				htp = strings.Trim(htp, " ")

				duplicatted := func(tpStr string) (b bool) {
					for _, t := range ht.TypeOrder {
						if t == tpStr {
							b = true
							return
						}
					}
					return
				}
				if !duplicatted(htp) {   // only store distinct types
					ht.TypeOrder = append(ht.TypeOrder, htp)
				}

				if _, ok := ht.Types[htp]; !ok {   // if initialed, do nothing
					ht.Types[htp] = make([]OneHeritage, 0)
				}

				s1.ChildrenFiltered("a").Each(func(i2 int, s2 *goquery.Selection) {  // a
					oh := OneHeritage{}
					oh.Country = worldHeritages.CountryList[worldHeritages.countryOrder[i]].Name
					oh.Href, _ = s2.Attr("href")
					oh.Href    = urlPrefix + strings.Trim(oh.Href, " ")
					oh.Name    = s2.Text()
					oh.Name    = strings.Trim(oh.Name, " ")

					//lg.Debugf("len(TypeOrder): %v, i3: %v", len(ht.TypeOrder), i3)
					ht.Types[htp] = append(ht.Types[htp], oh)
					//lg.Debugf("(%v) type: '%15v', href: '%v', text: '%50v', len: %v", i2, htp, oh.Href, oh.Name, len(ht.Types[htp]))
					//if oh.Href != "" && !strings.Contains(oh.Href, "transboundary") && !strings.Contains(oh.Href, "criteria_revision") {
					//	wohelist <- msg{Url: oh.Href, Status: false, Name: oh.Name}
					//	if oh.Href == "" {
					//		lg.Debugf("oh.Href is null")
					//	}
					//}
				})
			})
			//lg.Debugf("HeritageItem(%v): %v", i, ht)
			worldHeritages.CountryList[worldHeritages.countryOrder[i]].HeritageList = append(worldHeritages.CountryList[worldHeritages.countryOrder[i]].HeritageList, ht)
		})

		lg.Debugf("len CountryList: %v", len(worldHeritages.CountryList))
		//lg.Debugf("CountryList: %v", whl.CountryList[1])

		utils.Write2JsonFile(worldHeritages, "worldHeritageList.json")
		//wohelist <- msg{Status: true}
		//close(wohelist)
	})

	c.Visit("http://whc.unesco.org/en/list/")
}

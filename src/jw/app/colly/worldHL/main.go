package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/davyxu/golog"
	"github.com/gocolly/colly/v2"
)

var (
	urlPrefix = "http://whc.unesco.org"

)

func main() {
	lg := golog.New("world-heritage-list")
	c := colly.NewCollector(
		 colly.CacheDir("./whl"),
		)

	c.OnRequest(func(req *colly.Request) {
		//lg.Debugf("request: %v", req.URL)
	})

	c.OnResponse(func(rsp *colly.Response) {
		lg.Debugf("response...")
	})

	var whl WorldHeritageList
	whl.CountryList = make([]CountryItem, 0)

	c.OnHTML("#acc", func(e *colly.HTMLElement) {  // div
		e.DOM.ChildrenFiltered("h4").Each(func(i int, s *goquery.Selection) {  // h4

			ci := CountryItem{}
			href, _ := s.Children().Attr("href")
			ci.Href = urlPrefix + href // e.ChildAttr("a", "href")
			ci.Name = s.Children().Text()
			ci.Type, _ = s.Attr("id")  //e.Attr("id")

			whl.CountryList = append(whl.CountryList, ci)
			lg.Debugf("CountryItem: %v", ci)
		})

		e.DOM.ChildrenFiltered("div.list_site").Each(func(i int, s *goquery.Selection) { // div
			//lg.Debugf("div.list_site name: %v", i)
			//ht  := HeritageItem{}
			//e.DOM.Children().Each(func(i1 int, s1 *goquery.Selection) {   // ul
			//	s1.Children().Each(func(i2 int, s2 *goquery.Selection) {   // li
			//		htp, _ := s2.Attr("class")
			//		htp = strings.Trim(htp, " ")
			//		ht.TypeOrder = append(ht.TypeOrder, htp)
			//		ht.Types = make(map[string][]OneHeritage)
			//		ht.Types[htp] = make([]OneHeritage, 0)
			//		s2.Children().Each(func(i3 int, s3 *goquery.Selection) {  // a
			//			oh := OneHeritage{}
			//			oh.Href, _ = s3.Attr("href")
			//			oh.Href = strings.Trim(oh.Href, " ")
			//			oh.Name    = s3.Text()
			//			oh.Name    = strings.Trim(oh.Name, " ")
			//
			//			lg.Debugf("len(TypeOrder): %v, i3: %v", len(ht.TypeOrder), i3)
			//			ht.Types[htp] = append(ht.Types[htp], oh)
			//			lg.Debugf("type: '%15v', href: '%v', text: '%50v', len: %v", htp, oh.Href, oh.Name, len(ht.Types[htp]))
			//		})
			//	})
			//})
			//
			//whl.CountryList[i].HeritageList = append(whl.CountryList[i].HeritageList, ht)
		})

		//lg.Debugf("len CountryList: %v", len(whl.CountryList))
		//lg.Debugf("CountryList: %v", whl.CountryList)
		//for idx, _ := range whl.CountryList {
		//	if idx > 3 {
		//		break
		//	}
		//	lg.Debugf("CountryList: %v", whl.CountryList[idx])
		//}

	})

	c.Visit("http://whc.unesco.org/en/list/")

}

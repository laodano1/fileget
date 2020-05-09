package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/davyxu/golog"
	"github.com/gocolly/colly/v2"
	"strings"
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
		lg.Debugf("request: %v", req.URL)
	})

	c.OnResponse(func(rsp *colly.Response) {
		lg.Debugf("response...")
	})

	//h4Map := make(map[int]*colly.HTMLElement)
	//divMap := make(map[int]*colly.HTMLElement)
	//h4Map := make(map[int]string)
	//divMap := make(map[int]string)

	var whl WorldHeritageList
	whl.CountryList = make([]CountryItem, 0)

	c.OnHTML("#acc", func(e *colly.HTMLElement) {

		e.ForEach("h4", func(i int, e *colly.HTMLElement) {
			ci := CountryItem{}
			ci.Href = urlPrefix + e.ChildAttr("a", "href")
			ci.Name = e.ChildText("a")
			ci.Type = e.Attr("id")

			whl.CountryList = append(whl.CountryList, ci)
		})

		e.ForEach("div.list_site", func(i int, e *colly.HTMLElement) {   // div

			ht  := HeritageItem{}
			e.DOM.Children().Each(func(i1 int, s1 *goquery.Selection) {   // ul
				s1.Children().Each(func(i2 int, s2 *goquery.Selection) {   // li
					htp, _ := s2.Attr("class")
					htp = strings.Trim(htp, " ")
					ht.TypeOrder = append(ht.TypeOrder, htp)
					ht.Types = make(map[string][]OneHeritage)
					ht.Types[htp] = make([]OneHeritage, 0)
					s2.Children().Each(func(i3 int, s3 *goquery.Selection) {  // a
						oh := OneHeritage{}
						oh.Href, _ = s3.Attr("href")
						oh.Href = strings.Trim(oh.Href, " ")
						oh.Name    = s3.Text()
						oh.Name    = strings.Trim(oh.Name, " ")

						ht.Types[htp] = append(ht.Types[htp], oh)
						lg.Debugf("type: '%15v', href: '%v', text: '%50v', len: %v", htp, oh.Href, oh.Name, len(ht.Types[htp]))
					})
				})
			})

			whl.CountryList[i].HeritageList = append(whl.CountryList[i].HeritageList, ht)

		})

		lg.Debugf("CountryList: %v", whl.CountryList[0])
		lg.Debugf("CountryList: %v", whl.CountryList[1])
	})




	c.Visit("http://whc.unesco.org/en/list/")

}

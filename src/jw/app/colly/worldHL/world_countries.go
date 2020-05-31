package main

import (
	"fileget/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

var (
	unDataLPrefix = "http://data.un.org"
)

func getCountries() {
	c := colly.NewCollector(
			colly.CacheDir("./whl"),
			colly.UserAgent(UserAgent),
	)

	lg.Debugf("user agent: %v", c.UserAgent)

	c.OnHTML("#myUL", func(e *colly.HTMLElement) {  // ul

		e.DOM.Children().Each(func(i int, s *goquery.Selection) {  // li
			oc := OneCountry{}
			s.Find("td").Each(func(i1 int, s1 *goquery.Selection) {

				switch i1 {
				case 1:
					s1.ChildrenFiltered("img").Each(func(i2 int, s2 *goquery.Selection) {  // img
						oc.FlagHref, _ = s2.Attr("src")
						oc.FlagHref = strings.Replace(oc.FlagHref, "..", unDataLPrefix, 1)
						//util.Lg.Debugf("flag: %v", oc.FlagHref)
					})
				case 3:
					oc.Name   = s1.Text()
					oc.Region = s1.Find("font").Text()
					oc.Name   = strings.ReplaceAll(oc.Name, oc.Region, "")
				default:

				}
			})
			un.CountryList = append(un.CountryList, oc)
		})

		util.Lg.Debugf("un: %v", un)
		fName := fmt.Sprintf("%v%cctr%cunited_nations.json", exeDirPath, os.PathSeparator, os.PathSeparator)
		util.Write2JsonFile(un, fName)
	})

	c.Visit("http://data.un.org/en/index.html")

}
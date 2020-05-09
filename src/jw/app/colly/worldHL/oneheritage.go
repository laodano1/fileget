package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func GetHeritageInfo(usrStr string) {
	c := colly.NewCollector(
		colly.CacheDir("./whl"),
	)
	c.UserAgent = UserAgent

	c.OnHTML(".alternate", func(e *colly.HTMLElement) {
		e.DOM.ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
			lg.Debugf("parse heritage item page")

		})
	})

	c.Visit(usrStr)
}

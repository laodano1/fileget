package main

import "github.com/gocolly/colly/v2"

func getCountries() {
	c := colly.NewCollector(
			colly.CacheDir("./whl"),
	)
	c.UserAgent = UserAgent
	lg.Debugf("user agent: %v", c.UserAgent)

	c.OnHTML("#myUL", func(e *colly.HTMLElement) {

	})

	c.Visit("http://data.un.org/en/index.html")

}
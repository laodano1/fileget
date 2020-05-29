package main

import "github.com/gocolly/colly/v2"

func main() {
	c := colly.NewCollector()

	c.Visit("http://10.0.0.200:32567/login")


}

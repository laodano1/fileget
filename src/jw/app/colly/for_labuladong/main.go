package main

import (
	"bytes"
	"fileget/util"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"

func main() {

	c := colly.NewCollector(
		colly.CacheDir("./labuladong"),
		colly.UserAgent(UserAgent),
	)

	//c.OnResponse(func(res *colly.Response) {
	//
	//})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		title := e.DOM.Children().Filter("head").Children().Filter("title").Text()
		title = strings.ReplaceAll(title, " ", "")
		//if err := e.Response.Save(fmt.Sprintf("%v.html", title)); err != nil {
		//	util.Lg.Debugf("save error: %v", err)
		//}

		//util.Lg.Debugf("%v", e.Text)

		bts := bytes.ReplaceAll(e.Response.Body, []byte("href=\"/algo"), []byte("href=\"https://labuladong.gitbook.io/algo"))
		//
		//err := ioutil.WriteFile(fmt.Sprintf("%v.html", title), bts, 0777)
		//if err != nil {
		//	util.Lg.Debugf("save error: %v", err)
		//}
		//util.Lg.Debugf("%v", string(bts))

		reg, _ := regexp.Compile("href=\"https://labuladong.gitbook.io/.*\"")
		all := reg.FindAllStringSubmatch(string(bts), -1)
		//hpl := reg.FindStringSubmatch()
		util.Lg.Debugf("%v", all)

	})

	c.Visit("https://labuladong.gitbook.io/algo/")

}

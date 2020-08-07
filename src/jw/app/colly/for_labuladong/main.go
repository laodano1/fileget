package main

import (
	"fileget/util"
	"github.com/gocolly/colly/v2"
	"net"
	"strings"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"

var linksMap = make(map[string]string)

func Traverse(urlStr string, c *colly.Collector) {
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

		//bts := bytes.ReplaceAll(e.Response.Body, []byte("href=\"/algo"), []byte("href=\"https://labuladong.gitbook.io/algo"))
		//err := ioutil.WriteFile(fmt.Sprintf("%v.html", title), bts, 0777)
		//if err != nil {
		//	util.Lg.Debugf("save error: %v", err)
		//}

		e.ForEach("a", func(i int, e *colly.HTMLElement) {

			if !strings.HasPrefix(e.Attr("href"), "http") {
				txt := strings.TrimLeft(e.Text, " ")
				txt = strings.TrimRight(txt, " ")
				linksMap[strings.Replace(e.Attr("href"), "/algo", "https://labuladong.gitbook.io/algo", 1)] = txt
			}
			//util.Lg.Debugf("%v, %v", e.Text, e.Attr("href"))
		})

		for k, v := range linksMap {
			util.Lg.Debugf("%-16v, %v", v, k)
		}

	})

	c.Visit(urlStr)
}

func main() {
	urlStr := "https://labuladong.gitbook.io/algo/"
	c := colly.NewCollector(
		colly.CacheDir("./labuladong"),
		colly.UserAgent(UserAgent),
	)
	Traverse(urlStr, c)

}

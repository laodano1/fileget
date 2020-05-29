package main

import (
	"fileget/util"
	"gopkg.in/headzoo/surf.v1"
)


var (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
)

//func main() {
//
//
//	//c.Visit("http://10.0.0.200:32567/login")
//
//}

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("http://10.0.0.200:32567/login")
	if err != nil {
		panic(err)
	}
	html, _ := bow.Dom().Html()
	util.Lg.Debugf("title: %v", html)
	// Outputs: "The Go Programming Language"

	ret, _ := bow.Find("#pane-sa").Html()


	util.Lg.Debugf("ret: %v", ret)
}
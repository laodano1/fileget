package main

import (
	"fileget/util"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

const (
	i = 1 << iota // 1 << 0
	j = 2 << iota // 1 << 1
	k             // 2
	m = "aa"
	l = 2 << iota
)

func main2() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

func main() {

	var port string
	flag.StringVar(&port, "p", ":9999", "listen port setting")
	flag.Parse()

	e := gin.Default()

	e.GET("/sts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "zxs OK"})
	})

	if err := e.Run(port); err != nil {
		util.Lg.Errorf("server error: %v", err)
	}

}

func main1() {
	barrel := make(map[[2]int]bool)

	a := [2]int{1, 3}
	b := [2]int{2, 2}
	barrel[a] = true
	barrel[b] = true

	util.Lg.Debugf("map: %v", barrel)

	//bow := surf.NewBrowser()
	//err := bow.Open("http://10.0.0.200:32567/login")
	//if err != nil {
	//	panic(err)
	//}
	//html, _ := bow.Dom().Html()
	//util.Lg.Debugf("title: %v", html)
	//// Outputs: "The Go Programming Language"
	//
	//ret, _ := bow.Find("#pane-sa").Html()
	//
	//
	//util.Lg.Debugf("ret: %v", ret)
}

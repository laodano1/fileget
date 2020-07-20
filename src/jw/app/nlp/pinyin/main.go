package main

import (
	"fileget/src/jw/app/nlp/pinyin/lib"
	"fileget/util"
	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"net/http"
	"sync"
	"unicode"
)

var (
	Addr    = ":9999"
	pinyins sync.Map
)

//type rspJson struct {
//	py []string  `json:"py"`
//	hz []string  `json:"hz"`
//}

func StartServer() {
	gin.ForceConsoleColor()
	e := gin.Default()

	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	e.GET("/", func(c *gin.Context) {
		hans := c.Query("py")

		if hans == "" {
			c.JSON(http.StatusOK, "empty query")
		} else {
			rs := []rune(hans)
			rsp := new(lib.RspJsonDefault)
			rsp.Hz = make([]string, 0)
			rsp.Py = make([]string, 0)
			for i := range rs {
				var tmpItem string
				if unicode.Is(unicode.Han, rs[i]) { // 只处理汉字
					val, ok := pinyins.Load(string(rs[i]))
					if !ok {
						a := pinyin.NewArgs()
						a.Style = pinyin.Tone

						ps := pinyin.Pinyin(string(rs[i]), a)
						pinyins.Store(string(rs[i]), ps[0][0]) // cached in sync.Map

						tmpItem = ps[0][0]

					} else {
						tmpItem = val.(string)
					}
					rsp.Py = append(rsp.Py, tmpItem)
				} else {
					rsp.Py = append(rsp.Py, string(rs[i]))
				}
				rsp.Hz = append(rsp.Hz, string(rs[i]))
			}

			c.IndentedJSON(http.StatusOK, rsp)
			//c.JSON(http.StatusOK, rsp)
		}

	})

	if err := e.Run(Addr); err != nil {
		util.Lg.Errorf("server run failed: %v", err)
	}

}

func main() {
	//util.Lg.Debugf("pin yin: %v", pinyin.Pinyin(hans, a))
	StartServer()

}

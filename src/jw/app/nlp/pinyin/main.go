package main

import (
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

		//hans := "弍是"
		if hans == "" {
			c.JSON(http.StatusOK, "empty query")
		} else {
			var tmpHans []rune
			rs := []rune(hans)
			for i := range rs {
				if unicode.Is(unicode.Han, rs[i]) { // 只处理汉字
					tmpHans = append(tmpHans, rs[i])
				}
			}

			if len(tmpHans) == 0 {
				util.Lg.Debugf("no han character!")
				c.JSON(http.StatusOK, "no han character!")
				return
			}

			//outputHans := make([]string, 0)
			//opHans := make([][]string, 0)
			rsp := gin.H{}
			rsp["hz"] = make([]string, 0, len(tmpHans))
			rsp["py"] = make([]string, 0, len(tmpHans))
			for i := range tmpHans {

				val, ok := pinyins.Load(string(tmpHans[i]))
				if !ok {
					a := pinyin.NewArgs()
					a.Style = pinyin.Tone

					ps := pinyin.Pinyin(string(tmpHans[i]), a)
					pinyins.Store(string(tmpHans[i]), ps[0][0]) // cached in sync.Map

					//util.Lg.Debugf("pin yin: %v", ps)
					//outputHans = append(outputHans, ps[0][0])
					//opHans = append(opHans, ps[0])
					rsp["py"] = append(rsp["py"].([]string), ps[0][0])
				} else {
					//outputHans = append(outputHans, val.(string))
					//opHans = append(opHans,  []string{val.(string)})
					rsp["py"] = append(rsp["py"].([]string), val.(string))
				}
				//tmpItem := make([]string, 0)
				//tmpItem = append(tmpItem, string(tmpHans[i]))
				rsp["hz"] = append(rsp["hz"].([]string), string(tmpHans[i]))
			}

			//rsp["py"] = opHans


			//c.IndentedJSON(http.StatusOK, rsp)
			c.JSON(http.StatusOK, rsp)
		}

	})

	if err := e.Run(Addr); err != nil {
		util.Lg.Errorf("server run failed: %v", err)
	}

}

func main() {
	//hans := "范蠡"
	//hans := "弍是"
	//a := pinyin.NewArgs()
	//a.Style = pinyin.Tone
	//
	//util.Lg.Debugf("pin yin: %v", pinyin.Pinyin(hans, a))
	StartServer()

}

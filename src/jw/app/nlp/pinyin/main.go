package main

import (
	"fileget/util"
	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"net/http"
	"sync"
)

var (
	Addr    = ":9999"
	pinyins sync.Map
)

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
			val, ok := pinyins.Load(hans)
			if !ok {
				a := pinyin.NewArgs()
				a.Style = pinyin.Tone

				//for _, i := range hans {
				//	py := pinyin.Pinyin(string(i), a)
				//	pinyins.Store(string(i), py[0][0])
				//}
				ps := pinyin.Pinyin(hans, a)
				pinyins.Store(hans, ps)

				util.Lg.Debugf("pin yin: %v", ps)
				c.JSON(http.StatusOK, ps)
			} else {
				c.JSON(http.StatusOK, val)
			}

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

package main

import (
	"fileget/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
	"strings"
)


var (
	PORT = ":10000"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			util.Lg.Errorf("%v", e)
		}
	}()
	exeDir, _ := util.GetFullPathDir()
	fullPth := fmt.Sprintf("%v%cindex.html", exeDir, os.PathSeparator)
	util.Lg.Debugf("exe dir: %v", fmt.Sprintf("%v%cjs", exeDir, os.PathSeparator))
	gin.ForceConsoleColor()
	e := gin.Default()

	e.Static("/js", fmt.Sprintf("%v%cjs", exeDir, os.PathSeparator))
	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	e.GET("/", func(c *gin.Context) {
		tp := template.Must(template.New("barchart").Funcs(template.FuncMap{
			"toLower": func(name string) string {
				return strings.ToLower(name)
			},
		}).ParseFiles(fullPth))

		dt := gin.H{
			"title": "echarts sample",
		}

		rd := render.HTML{
			Template: tp,
			Name:     "barchart",
			Data:     dt,
		}

		c.Render(http.StatusOK, rd)
	})

	if err := e.Run(PORT); err != nil {
		util.Lg.Errorf("run gin failed: %v", err)
	}
}

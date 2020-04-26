package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
)

func homepage(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%cindex.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("homepage").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "my website",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "homepage",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func buttons(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%clyear_ui_buttons.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	//lg.Debugf("button tmpl: %v", fullPth)
	tp := template.Must(template.New("button").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "buttons",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "button",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func (m *myBackend) addRoutes() {
	m.e.GET("/", homepage)

	uierg := m.e.Group("/uie")

	m.e.Static("/uie/css",    fmt.Sprintf("%v%ctemplate%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static("/uie/js",     fmt.Sprintf("%v%ctemplate%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static("/uie/images", fmt.Sprintf("%v%ctemplate%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static("/uie/fonts",  fmt.Sprintf("%v%ctemplate%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
	uierg.Handle("GET", "/button", buttons)

	//
	//m.e.GET("/ui/:name", idx)
}





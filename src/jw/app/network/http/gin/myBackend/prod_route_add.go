package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
)

func prodHP(c *gin.Context)  {
	fullPth := fmt.Sprintf("%v%ctmpl%cindex.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("homepage").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "production website",
		"radio" : "Radio",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "homepage",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)

}

func (m *myBackend) addProductRoutes() {
	m.e.GET("/", prodHP)
}

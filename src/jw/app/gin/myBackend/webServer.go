package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func NewBK() *myBackend {
	gin.ForceConsoleColor()
	e := gin.Default()
	e.Static("/css", "./template/css")
	e.Static("/js", "./template/js")
	e.Static("/images", "./template/images")
	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	return &myBackend{
		name: "my backend",
		e:     e,
	}
}

func (m *myBackend) StartBK(add string) (err error) {
	err = m.e.Run(add)
	if err != nil {
		lg.Errorf("backend server run failed: %v", err)
		return
	}
	return

}

func homepage(c *gin.Context) {

	exeAbsPath, err := os.Executable()
	if err != nil {
		lg.Errorf("get os.Executable failed: %v", err)
	}
	//


	// Working Directory
	//wd, err := os.Getwd()
	//if err != nil {
	//	lg.Errorf("%v", err)
	//}
	lg.Debugf("os.Executable path: %v", exeAbsPath)

	dir := filepath.Dir(exeAbsPath)
	fullPth := fmt.Sprintf("%v%ctemplate\\index.html", dir, os.PathSeparator)
	lg.Debugf("template path: %v", fullPth)

	//template.ParseFiles("index.html")

	tp := template.Must(template.New("homepage").ParseFiles(fullPth))
	//if err != nil {
	//	lg.Errorf("parse file failed: %v", err)
	//}
	//
	dt := &gin.H{
		"title" : "hello world",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "homepage",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func idx(c *gin.Context) {
	filename := c.Param("name")
	lg.Debugf("in idx handler. filename: %v.html", filename)

	c.HTML(http.StatusOK, "index.html", nil)
}

func (m *myBackend) addRoutes() {
	m.e.GET("/", homepage)
	//m.e.GET("/tp/:name", idx)
}




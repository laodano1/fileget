package main

import (
	"fileget/src/jw/app/gin/myBackend/utis"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
)

func NewBK() (*myBackend, error) {
	gin.ForceConsoleColor()
	e := gin.Default()
	//var err error
	exeAbsPath, _ = utis.GetFullPathDir()

	// set static files
	e.Static("/css", fmt.Sprintf("%v%ctemplate%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/js", fmt.Sprintf("%v%ctemplate%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/images", fmt.Sprintf("%v%ctemplate%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/fonts", fmt.Sprintf("%v%ctemplate%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))

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
	}, nil
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
	//var err error
	//exeAbsPath, err = utis.GetFullPathDir()
	//if err != nil {
	//	lg.Errorf("%v", err)
	//	c.JSON(http.StatusInternalServerError, &gin.H{"result": "failed", "description": err.Error()})
	//}


	fullPth := fmt.Sprintf("%v%ctemplate\\index.html", exeAbsPath, os.PathSeparator)
	//lg.Debugf("template path: %v", fullPth)


	tp := template.Must(template.New("homepage").ParseFiles(fullPth))

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




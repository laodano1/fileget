package main

import (
	"encoding/json"
	"fileget/src/jw/app/gin/world_heritage/utils"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	port = ":10000"
	exeAbsPath string
	lg = golog.New("world_heritage")
)

func NewGinServer() *Myserver {
	gin.ForceConsoleColor()
	e := gin.Default()
	//var err error
	exeAbsPath, _ = utils.GetFullPathDir()

	// set static files
	e.Static("/css",    fmt.Sprintf("%v%cpublic%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/js", fmt.Sprintf("%v%cpublic%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/custom_js", fmt.Sprintf("%v%cpublic%ccustom_js", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/assets", fmt.Sprintf("%v%cpublic%cassets", exeAbsPath, os.PathSeparator, os.PathSeparator))
	//e.Static("/fonts",  fmt.Sprintf("%v%ctmpl%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
	//e.Static("/video",  fmt.Sprintf("%v%ctmpl%cvideo", exeAbsPath, os.PathSeparator, os.PathSeparator))

	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	return &Myserver{
		e: e,
	}
}


func (ms *Myserver) hp(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctmpl%cinfo.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("info").Funcs(template.FuncMap{
		"toLower": func(name string) string {
			return strings.ToLower(name)
		},
	}).ParseFiles(fullPth))

	dt := gin.H{
		"title": "haha",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "info",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)

}

func (ms *Myserver) lpjson(c *gin.Context) {
	jsonb, err := ioutil.ReadFile(exeAbsPath + "/public/lpinfors.json")
	if err != nil {
		lg.Errorf("read json file failed: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": "failed", "description": err})
		return
	}

	allLP := new(LpList)
	err = json.Unmarshal(jsonb, allLP)
	if err != nil {
		lg.Errorf("unmarshal failed: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": "failed", "description": err})
		return
	}
	lg.Debugf("%v", allLP)
	c.IndentedJSON(http.StatusOK, allLP)
}

func (ms *Myserver) AddRoutes() {
	ms.e.GET("/", ms.hp)
	ms.e.GET("/lpjson", ms.lpjson)
}

func (ms *Myserver) Start() {
	if err := ms.e.Run(port); err != nil {
		lg.Errorf("start gin server failed: %v", err)
	}
}



package main

import (
	"fileget/src/jw/app/gin/myBackend/utis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var sbr = `
{
  "list": [{
	"name": "Media", 
	"subItems": [
		{"name":"mp3", "href": "/mp3"}, 
		{"name": "mp4", "href": "/mp4"}, 
		{"name": "mkv", "href": "/mkv"}
	],
    "pageCnt": [
        {"pageObjs": []},
        {"pageObjs": [{"name": "tesla", "href": "/video/tesla.mp4"}] },
        {"pageObjs": [
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"},
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"},
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"},
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"},
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"},
			{"name": "ye-wen-4", "href": "/video/ye-wen-4.mkv"}

		]}	
     ] 
   }]
}
`
var sidebarData sidebar


func NewBK() (*myBackend, error) {
	//if err := getSBData(); err != nil {
	//	os.Exit(1)
	//}

	gin.ForceConsoleColor()
	e := gin.Default()
	//var err error
	exeAbsPath, _ = utis.GetFullPathDir()

	// set static files
	e.Static("/css",    fmt.Sprintf("%v%ctmpl%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/js",     fmt.Sprintf("%v%ctmpl%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/images", fmt.Sprintf("%v%ctmpl%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/fonts",  fmt.Sprintf("%v%ctmpl%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
	e.Static("/video",  fmt.Sprintf("%v%ctmpl%cvideo", exeAbsPath, os.PathSeparator, os.PathSeparator))

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







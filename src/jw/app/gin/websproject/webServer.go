package main

import (
	"encoding/json"
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

func getSBData() (err error) {
	err = json.Unmarshal([]byte(sbr), &sidebarData)
	if err != nil {
		lg.Errorf("unmarshal sidebar data failed: %v", err)
		return
	}
	lg.Debugf("sidebar: %v", sidebarData)
	return
}

func LoadData() {
	subIList := make([]sbSubItem, 0)
	subIList = append(subIList, sbSubItem{Name: "mp3", Href: "/mp3"})
	subIList = append(subIList, sbSubItem{Name: "mp4", Href: "/mp4"})
	subIList = append(subIList, sbSubItem{Name: "mkv", Href: "/mkv"})

	pgItems1 := make([]pageItem, 0)
	pgItems1 = append(pgItems1, pageItem{})
	pgcnt1 := pageContent{PageObjs: pgItems1}

	pgItems2 := make([]pageItem, 0)
	pgItems2 = append(pgItems2, pageItem{Name: "tesla", Href: "/video/tesla.mp4"})
	pgcnt2 := pageContent{PageObjs: pgItems2}

	pgItems3 := make([]pageItem, 0)
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgcnt3 := pageContent{PageObjs: pgItems3}

	pgcList := make([]pageContent, 0)
	pgcList = append(pgcList, pgcnt1)
	pgcList = append(pgcList, pgcnt2)
	pgcList = append(pgcList, pgcnt3)

	sbmi := sidebarMainItem{
		Name:     "Media",
		SubItems: subIList,
		PageCnt:  pgcList,
	}

	sbmiList := make([]sidebarMainItem, 0)
	sbmiList = append(sbmiList, sbmi)
	sbObj := &sidebar{
		List: sbmiList,
	}

	sidebarData = *sbObj
}

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







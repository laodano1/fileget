package main

import (
	"encoding/json"
	"fileget/src/jw/app/gin/world_heritage/utils"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	port = ":10000"
	exeAbsPath string
	lg = golog.New("world_heritage")
	allJsonFilePath []string
	allHtailData []HeritageDetail
)

func PreStartActions() (err error)  {
	lg.SetParts(golog.LogPart_Level, golog.LogPart_TimeMS, golog.LogPart_ShortFileName)
	lg.EnableColor(true)
	startTime := time.Now()

	exeAbsPath, _ = utils.GetFullPathDir()
	allJsonFilePath, err = utils.GetFiles(fmt.Sprintf("%v%ctmp", exeAbsPath, os.PathSeparator), "json")
	if err != nil {
		lg.Errorf("%v", err)
	}

	allHtailData, err = getJsonDataList(allJsonFilePath, reflect.TypeOf(HeritageDetail{}))
	if err != nil {
		lg.Errorf("%v", err)
	}


	lg.Debugf("*********************** preparing spent: %v seconds.", time.Now().Sub(startTime))
	return
}

func NewGinServer() *Myserver {

	gin.ForceConsoleColor()
	e := gin.Default()

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


func (ms *Myserver) hp(tmpl string) func(c *gin.Context) {

	return func(c *gin.Context) {
		fullPth := fmt.Sprintf("%v%ctmpl%c%v.html", exeAbsPath, os.PathSeparator, os.PathSeparator, tmpl)
		tp := template.Must(template.New(tmpl).Funcs(template.FuncMap{
			"toLower": func(name string) string {
				return strings.ToLower(name)
			},
		}).ParseFiles(fullPth))

		dt := gin.H{
			"title": "世界遗产坐标",
		}

		rd := render.HTML{
			Template: tp,
			Name:     tmpl,
			Data:     dt,
		}

		c.Render(http.StatusOK, rd)
	}
}


func (ms *Myserver) getInfo(data []HeritageDetail, infoType string) func(c *gin.Context) {

	return func (c *gin.Context) {
		//var mth string
		switch infoType {
		case info_type_WHL:
			//jsonb, err := utils.ReadWHLJson(fileNames[0])
			//if err != nil {
			//	lg.Errorf("%v", err)
			//	c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err})
			//	return
			//}
			//allH := new(HeritageDetail)
			//err = json.Unmarshal(jsonb, allH)
			//if err != nil {
			//	c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err})
			//	return
			//}
			c.IndentedJSON(http.StatusOK, data)
		case info_type_Loupan :
			//mth = c.Query("month")
			//lg.Debugf("month: %v", mth)
			//jsonb, err := utils.ReadJson(exeAbsPath, fileNames[0])
			//if err != nil {
			//	lg.Errorf("%v", err)
			//	c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err})
			//	return
			//}
			//allLP := new(LpList)
			//err = json.Unmarshal(jsonb, allLP)
			//if err != nil {
			//	c.IndentedJSON(http.StatusInternalServerError, gin.H{"result": err})
			//	return
			//}
			//c.IndentedJSON(http.StatusOK, allLP.Month[mth])
		default:
			c.IndentedJSON(http.StatusOK, gin.H{"result": "unknown type!"})
		}
    }
}

func (ms *Myserver) AddRoutes() {
	ms.e.GET("/", ms.hp("whl"))
	//ms.e.GET("/lpjson", ms.getInfo([]string{json_loupan}, info_type_Loupan))

	//ms.e.GET("/whl",   ms.hp("whl"))
	ms.e.GET("/whldt", ms.getInfo(allHtailData, info_type_WHL))
}

func (ms *Myserver) Start() {
	if err := ms.e.Run(port); err != nil {
		lg.Errorf("start gin server failed: %v", err)
	}
}

func getJsonDataList(files []string, p reflect.Type) (list []HeritageDetail, err error) {
	for _, f := range files {
		var jsonb []byte
		jsonb, err = utils.ReadWHLJson(f)
		if err != nil {
			return nil, err
		}
		hd := new(HeritageDetail)
		json.Unmarshal(jsonb, hd)
		if err != nil {
			return nil, err
		}
		list = append(list, *hd)
	}

	return
}


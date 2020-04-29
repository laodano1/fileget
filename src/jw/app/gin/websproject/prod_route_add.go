package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func generalHandle(tmplName string, midx, sidx int) func(c *gin.Context)  {
	return func(c *gin.Context) {
		fullPth := fmt.Sprintf("%v%ctmpl%c%v.html", exeAbsPath, os.PathSeparator, os.PathSeparator, tmplName)
		tp := template.Must(template.New(tmplName).Funcs(template.FuncMap{
			"toLower": func(name string) string {
				return strings.ToLower(name)
			},
			"isEmpty": func(list []pageItem) bool {
				if len(list) > 0 {
					//lg.Debugf("not empty")
					return false
				}
				lg.Debugf("it's empty")
				return true
			},
			"isZero": func(id int) bool {
				if id == 0 {
					return true
				}
				return false
			},
		}).ParseFiles(fullPth))
		//
		dt := &gin.H{
			"title" : tmplName,
			"radio": strings.Split(tmplName, "-")[1] ,
			"sidebarlists" : sidebarData.List,
			"pagecontent"  : sidebarData.List[midx].PageCnt[sidx].PageObjs,
		}

		//lg.Debugf("req url: %v. page content objects: %v", c.Request.URL, sidebarData.List[midx].PageCnt[sidx].PageObjs)

		rd := render.HTML{
			Template: tp,
			Name:     tmplName,
			Data:     dt,
		}

		c.Render(http.StatusOK, rd)
	}
}


func hp(c *gin.Context)  {
	fullPth := fmt.Sprintf("%v%ctmpl%cindex.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("homepage").Funcs(template.FuncMap{
		"toLower": func(name string) string {
			return strings.ToLower(name)
		},
	}).ParseFiles(fullPth))
	//

	dt := &gin.H{
		"title" : "production website",
		"sidebarlists" : sidebarData.List,

	}

	rd := render.HTML{
		Template: tp,
		Name:     "homepage",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func cfg(c *gin.Context)  {
	fullPth := fmt.Sprintf("%v%ctmpl%cweb-config.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("config").Funcs(template.FuncMap{
		"toLower": func(name string) string {
			return strings.ToLower(name)
		},
	}).ParseFiles(fullPth))
	//

	dt := &gin.H{
		"title" : "production website",
		//"sidebarlists" : sidebarData.List,

	}

	rd := render.HTML{
		Template: tp,
		Name:     "config",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func (m *myBackend) addProductRoutes() {
	m.e.GET("/", hp)
	m.e.GET("/cfg", cfg)


	for idx, dirItem := range staticDir {
		m.e.Static(fmt.Sprintf("/%v", dirItem),   dirList.DirList[idx])
	}

	for midx, sbitem := range sidebarData.List {
		mdg := m.e.Group(fmt.Sprintf("/%v", strings.ToLower(sbitem.Name)))
		m.staticHandle(fmt.Sprintf("/%v", strings.ToLower(sbitem.Name)))
		for sidx, subItem := range sbitem.SubItems {
			tmp := fmt.Sprintf("%v-%v", strings.ToLower(sbitem.Name), subItem.Name)
			//lg.Debugf("tmpl: %v", tmp)
			mdg.GET(fmt.Sprintf("/%v", subItem.Name), generalHandle(tmp, midx, sidx))
		}
	}

}

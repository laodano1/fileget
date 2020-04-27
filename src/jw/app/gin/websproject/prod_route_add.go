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

func generalHandle(tmplName string) func(c *gin.Context)  {
	return func(c *gin.Context) {
		fullPth := fmt.Sprintf("%v%ctmpl%c%v.html", exeAbsPath, os.PathSeparator, os.PathSeparator, tmplName)
		tp := template.Must(template.New(tmplName).Funcs(template.FuncMap{
			"tolowwer": func(name string) string {
				return strings.ToLower(name)
			},
		}).ParseFiles(fullPth))
		//

		dt := &gin.H{
			"title" : tmplName,
			"sidebarlists" : sidebarData.List,

		}

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
		"tolowwer": func(name string) string {
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

func (m *myBackend) addProductRoutes() {
	m.e.GET("/", hp)

	for _, sbitem := range sidebarData.List {
		mdg := m.e.Group(fmt.Sprintf("/%v", strings.ToLower(sbitem.Name)))
		m.staticHandle(fmt.Sprintf("/%v", strings.ToLower(sbitem.Name)))
		for _, subItem := range sbitem.SubItems {
			tmp := fmt.Sprintf("%v-%v", strings.ToLower(sbitem.Name), subItem.Name)
			lg.Debugf("tmpl: %v", tmp)
			mdg.GET(fmt.Sprintf("/%v", subItem.Name), generalHandle(tmp))
		}
	}

}

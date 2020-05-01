package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
	"os/exec"
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

func (m *myBackend)  staticHandle(group string)  {
	m.e.Static(fmt.Sprintf("%v/css", group),    fmt.Sprintf("%v%ctmpl%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/js", group),     fmt.Sprintf("%v%ctmpl%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/images", group), fmt.Sprintf("%v%ctmpl%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/fonts", group),  fmt.Sprintf("%v%ctmpl%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/video", group),  fmt.Sprintf("%v%ctmpl%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
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

func forshutdown(c *gin.Context)  {
	fullPth := fmt.Sprintf("%v%ctmpl%cshutdown.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("shutdown").Funcs(template.FuncMap{
		"toLower": func(name string) string {
			return strings.ToLower(name)
		},
	}).ParseFiles(fullPth))
	//

	dt := &gin.H{
		"title" : "shutdown raspberry pi 3 website",
		"radio" : "shutdown raspberry pi 3",

	}

	rd := render.HTML{
		Template: tp,
		Name:     "shutdown",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func shutdown(c *gin.Context) {
	codeStr := c.Param("code")
	lg.Debugf("code string: %v", codeStr)

	if codeStr == "911" {
		cmd := exec.Command("shutdown")
		err := cmd.Start()
		if err != nil {
			lg.Errorf("shutdown command executes failed!")
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"result" : "shutdown command executes failed!"})
			return
		}
		c.IndentedJSON(http.StatusOK, &gin.H{"result": "ok"})
	} else {
		c.IndentedJSON(http.StatusOK, &gin.H{ "result": "code invalid! not shutdown!"} )
	}
}

func (m *myBackend) addProductRoutes() {
	m.e.GET("/", hp)
	m.e.GET("/cfg", cfg)
	m.e.GET("/forsd", forshutdown)
	m.e.GET("/sd/:code", shutdown)


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

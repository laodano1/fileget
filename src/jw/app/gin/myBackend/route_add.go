package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"os"
)

func homepage(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%cindex.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	tp := template.Must(template.New("homepage").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "my website",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "homepage",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func buttons(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%clyear_ui_buttons.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	//lg.Debugf("button tmpl: %v", fullPth)
	tp := template.Must(template.New("button").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "buttons",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "button",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func belements(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%clyear_forms_elements.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	//lg.Debugf("button tmpl: %v", fullPth)
	tp := template.Must(template.New("elements").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "elements",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "elements",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func examples(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%clyear_pages_doc.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	//lg.Debugf("button tmpl: %v", fullPth)
	tp := template.Must(template.New("doclist").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "doc list",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "doclist",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func datechs(c *gin.Context) {
	fullPth := fmt.Sprintf("%v%ctemplate%clyear_js_datepicker.html", exeAbsPath, os.PathSeparator, os.PathSeparator)
	//lg.Debugf("button tmpl: %v", fullPth)
	tp := template.Must(template.New("dtch").ParseFiles(fullPth))
	//
	dt := &gin.H{
		"title" : "date choose",
	}

	rd := render.HTML{
		Template: tp,
		Name:     "dtch",
		Data:     dt,
	}

	c.Render(http.StatusOK, rd)
}

func (m *myBackend)  staticHandle(group string)  {
	m.e.Static(fmt.Sprintf("%v/css", group),    fmt.Sprintf("%v%ctemplate%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/js", group),     fmt.Sprintf("%v%ctemplate%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/images", group), fmt.Sprintf("%v%ctemplate%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/fonts", group),  fmt.Sprintf("%v%ctemplate%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
}

func (m *myBackend) addRoutes() {
	m.e.GET("/", homepage)

	uierg := m.e.Group("/uie")
	m.staticHandle("/uie")
	uierg.Handle("GET", "/button", buttons)

	formserg := m.e.Group("/forms")
	m.staticHandle("/forms")
	formserg.Handle("GET", "/elements", belements)

	exapserg := m.e.Group("/examples")
	m.staticHandle("/examples")
	exapserg.Handle("GET", "/doclist", examples)

	jsprg := m.e.Group("/jpin")
	m.staticHandle("/jpin")
	jsprg.Handle("GET", "/dtch", datechs)

}





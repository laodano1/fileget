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

func (m *myBackend)  staticHandle(group string)  {
	m.e.Static(fmt.Sprintf("%v/css", group),    fmt.Sprintf("%v%ctemplate%ccss", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/js", group),     fmt.Sprintf("%v%ctemplate%cjs", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/images", group), fmt.Sprintf("%v%ctemplate%cimages", exeAbsPath, os.PathSeparator, os.PathSeparator))
	m.e.Static(fmt.Sprintf("%v/fonts", group),  fmt.Sprintf("%v%ctemplate%cfonts", exeAbsPath, os.PathSeparator, os.PathSeparator))
}


func (m *myBackend) generalHandle(tmplFileName string) func(c *gin.Context) {

	return func(c *gin.Context) {
		fullPth := fmt.Sprintf("%v%ctemplate%c%v", exeAbsPath, os.PathSeparator, os.PathSeparator, tmplFileName)
		//lg.Debugf("button tmpl: %v", fullPth)

		nitem := strings.Split(strings.Split(tmplFileName, ".")[0], "_")
		tp := template.Must(template.New(fmt.Sprintf("%v_%v", nitem[1], nitem[2])).ParseFiles(fullPth))
		//
		dt := &gin.H{
			"title" : fmt.Sprintf("%v %v", nitem[1], nitem[2]),
		}

		rd := render.HTML{
			Template: tp,
			Name:     fmt.Sprintf("%v_%v", nitem[1], nitem[2]),
			Data:     dt,
		}

		c.Render(http.StatusOK, rd)
	}

}

func (m *myBackend) addRoutes() {
	m.e.GET("/", homepage)

	uierg := m.e.Group("/uie")
	m.staticHandle("/uie")
	uierg.Handle("GET", "/button",   m.generalHandle("lyear_ui_buttons.html"))
	uierg.Handle("GET", "/cards",    m.generalHandle("lyear_ui_cards.html"))
	uierg.Handle("GET", "/grids",    m.generalHandle("lyear_ui_grid.html"))
	uierg.Handle("GET", "/icons",    m.generalHandle("lyear_ui_icons.html"))
	uierg.Handle("GET", "/tables",   m.generalHandle("lyear_ui_tables.html"))
	uierg.Handle("GET", "/modals",   m.generalHandle("lyear_ui_modals.html"))
	uierg.Handle("GET", "/popover",  m.generalHandle("lyear_ui_tooltips_popover.html"))
	uierg.Handle("GET", "/alerts",     m.generalHandle("lyear_ui_alerts.html"))
	uierg.Handle("GET", "/pagination", m.generalHandle("lyear_ui_pagination.html"))
	uierg.Handle("GET", "/progress",   m.generalHandle("lyear_ui_progress.html"))
	uierg.Handle("GET", "/tabs",       m.generalHandle("lyear_ui_tabs.html"))
	uierg.Handle("GET", "/steps",      m.generalHandle("lyear_ui_step.html"))
	uierg.Handle("GET", "/typography", m.generalHandle("lyear_ui_typography.html"))
	uierg.Handle("GET", "/other",      m.generalHandle("lyear_ui_other.html"))

	formserg := m.e.Group("/forms")
	m.staticHandle("/forms")
	formserg.Handle("GET", "/elements", m.generalHandle("lyear_forms_elements.html"))
	formserg.Handle("GET", "/radio",    m.generalHandle("lyear_forms_radio.html"))
	formserg.Handle("GET", "/switch",   m.generalHandle("lyear_forms_switch.html"))
	formserg.Handle("GET", "/checkbox", m.generalHandle("lyear_forms_checkbox.html"))

	exapserg := m.e.Group("/pages")
	m.staticHandle("/pages")
	exapserg.Handle("GET", "/doclist", m.generalHandle("lyear_pages_doc.html"))
	exapserg.Handle("GET", "/gallery", m.generalHandle("lyear_pages_gallery.html"))
	exapserg.Handle("GET", "/config", m.generalHandle("lyear_pages_config.html"))
	exapserg.Handle("GET", "/rabc",   m.generalHandle("lyear_pages_rabc.html"))
	exapserg.Handle("GET", "/adddoc", m.generalHandle("lyear_pages_add_doc.html"))
	exapserg.Handle("GET", "/wizard", m.generalHandle("lyear_pages_guide.html"))
	exapserg.Handle("GET", "/lgin",   m.generalHandle("lyear_pages_login.html"))
	exapserg.Handle("GET", "/err",    m.generalHandle("lyear_pages_error.html"))

	jsprg := m.e.Group("/jpin")
	m.staticHandle("/jpin")
	jsprg.Handle("GET", "/dtpkr", m.generalHandle("lyear_js_datepicker.html"))
	jsprg.Handle("GET", "/sldr", m.generalHandle("lyear_js_sliders.html"))
	jsprg.Handle("GET", "/clpk", m.generalHandle("lyear_js_colorpicker.html"))
	jsprg.Handle("GET", "/chts", m.generalHandle("lyear_js_chartjs.html"))
	jsprg.Handle("GET", "/jcf", m.generalHandle("lyear_js_jconfirm.html"))
	jsprg.Handle("GET", "/tg",  m.generalHandle("lyear_js_tags_input.html"))
	jsprg.Handle("GET", "/ntf", m.generalHandle("lyear_js_notify.html"))

}





package main

import "github.com/gin-gonic/gin"

type myBackend struct {
	e     *gin.Engine
	name  string
}

type sbSubItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type sidebarMainItem struct {
	Name     string      `json:"name"`
	SubItems []sbSubItem `json:"subItems"`
}

type sidebar struct {
	List  []sidebarMainItem `json:"list"`
}

type pageContent struct {

}




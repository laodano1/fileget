package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type myBackend struct {
	e     *gin.Engine
	name  string
}

type pageItem struct {
	Name string `json:"name"`
	Href  string   `json:"href"`
}

type pageContent struct {
	PageObjs  []pageItem `json:"pageObjs"`
}

type sbSubItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type sidebarMainItem struct {
	Name     string        `json:"name"`
	SubItems []sbSubItem   `json:"subItems"`
	PageCnt  []pageContent `json:"pageCnt"`
}

type sidebar struct {
	List  []sidebarMainItem `json:"list"`
}

type dbObj struct {
	Db *gorm.DB

}

type dirsJsonObj struct{
	DirList []string `json:"dirlist"`
}


type oneType struct {
	typeName string
	names  []string
	paths   []string
}


type filesList struct {
	list map[string]*oneType
}





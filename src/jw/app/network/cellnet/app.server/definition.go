package main

import "github.com/gin-gonic/gin"

type clientStatus struct {
	id    int
	phase string
}

type allClientStatus struct {
	cliNum int
	asMap map[int]*clientStatus
}

type chObj struct {
	stch chan *clientStatus
	acs  allClientStatus
}

type webSrvObj struct {
	name string
	e    *gin.Engine
	cor  *chObj
}

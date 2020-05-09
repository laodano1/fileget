package main

import "github.com/gin-gonic/gin"

type Myserver struct {
	e *gin.Engine
}

type LpItem struct{
	Name  string `json:"name"`
	Block string `json:"block"`
	Timestamp string   `json:"timestamp"`
	Subway    []string `json:"subway"`
	Xuequ     []string `json:"xuequ"`
	Type      string   `json:"type"`
	Unit      string   `json:"unit"`
	Price     string   `json:"price"`
	Coordinate []string `json:"coordinate"`
}


type LpList struct {
	Month map[string][]LpItem `json:"month"`
	//Month map[string][]LpItem

}
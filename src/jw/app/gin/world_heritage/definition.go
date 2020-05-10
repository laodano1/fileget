package main

import "github.com/gin-gonic/gin"

type Myserver struct {
	e *gin.Engine
}

type LpItem struct{
	Name  string `json:"name"`
	Block string `json:"block"`
	Timestamp string   `json:"timestamp"`
	subway    []string `json:"subway"`
	xuequ     []string `json:"xuequ"`
	Type      string   `json:"type"`
	Unit      string   `json:"unit"`
	Price     string   `json:"price"`
	Coordinate []string `json:"coordinate"`
}

type LpList struct {
	Month map[string][]LpItem `json:"month"`
}


type OneHeritage struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Country string `json:"country"`
}

type HeritageItem struct {
	TypeOrder []string  `json:"type_order"`  // type order to manage types map key sequence
	Types     map[string][]OneHeritage `json:"types"`
	Country   string `json:"country"`
}

type CountryItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Type string `json:"type,omitempty"`
	HeritageList []HeritageItem `json:"heritage_list"`
}

type WorldHeritageList struct {
	CountryList map[string]*CountryItem `json:"country_list"`
}
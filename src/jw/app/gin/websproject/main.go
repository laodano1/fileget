package main

import (
	"github.com/davyxu/golog"
)

const (
	add = ":10000"
)

var (
	lg = golog.New("my-backend")
	exeAbsPath string
	myDB *dbObj
)


func main() {
	var err error
	LoadData()

	bk, err := NewBK()
	if err != nil {
		lg.Errorf("%v", err)
	}
	//if mode == "book" {
	//	lg.Debugf("work mode: %v", mode)
	//	bk.addGuideLineRoutes()
	//} else {
		bk.addProductRoutes()
	//}

    bk.StartBK(add)
}

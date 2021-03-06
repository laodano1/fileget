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
	//LoadData()

	err = LoadMediaInfo()
	if err != nil {
		return
	}

	bk, err := NewBK()
	if err != nil {
		lg.Errorf("%v", err)
	}

	bk.addProductRoutes()

    bk.StartBK(add)
}

package main

import (
	"flag"
	"github.com/davyxu/golog"
)

const (
	add = ":9999"
)

var (
	lg = golog.New("my-backend")
	exeAbsPath string
)


func main() {
	var mode string
	flag.StringVar(&mode, "m", "book", "working mode. support: book | prod")
	//flag.StringVar(&mode, "m", "prod", "working mode. support: book | prod")
	flag.Parse()
	bk, err := NewBK()
	if err != nil {
		lg.Errorf("%v", err)
	}
	if mode == "book" {
		lg.Debugf("work mode: %v", mode)
		bk.addGuideLineRoutes()
	} else {
		bk.addProductRoutes()
	}

    bk.StartBK(add)
}

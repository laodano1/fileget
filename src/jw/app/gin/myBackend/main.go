package main

import "github.com/davyxu/golog"

const (
	add = ":9999"
)

var (
	lg = golog.New("my-backend")
	exeAbsPath string
)


func main() {
	bk, err := NewBK()
	if err != nil {
		lg.Errorf("%v", err)
	}
	bk.addRoutes()
    bk.StartBK(add)
}

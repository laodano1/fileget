package main

import "github.com/davyxu/golog"

const (
	add = ":9999"
)

var (
	lg = golog.New("my-backend")
)


func main() {
	bk := NewBK()
	bk.addRoutes()
    bk.StartBK(add)
}

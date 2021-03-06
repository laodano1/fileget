package main

import (
	"github.com/davyxu/golog"
)

var (
	srvAdd = "10.0.0.156"
)

var (
	lg   = golog.New("new adapter")
	PORT = "0.0.0.0:80"
)

func main() {
	lg.SetParts()
	adap := NewAdapter()

}

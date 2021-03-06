package main

import (
	"fileget/util"
	"github.com/davyxu/golog"
)

const (
	b = 0 - 1
)

type la struct {
	a string
}

type lai interface {
}

func aaa(i lai) {
	//lg.Debugf("i: %t", i.(type))
	switch ia := i.(type) {
	case la:
		lg.Debugf("la: %v", ia)
	}
}

var lg = golog.New("condition")

func main() {

	a := 0 == 0

	lg.Debugf("a: %v, b: %v", a, b)
	util.Logger.Sugar().Debugf("a: %v, b: %v", a, b)
	util.Logger.Sugar().Infof("a: %v, b: %v", a, b)

	t := 1 << 11 //
	lg.Debugf("t: %d", t)

	latmp := la{a: "test"}
	aaa(latmp)

	//net.Listen()

}

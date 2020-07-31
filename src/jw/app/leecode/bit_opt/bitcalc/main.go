package main

import (
	"fileget/util"
)

func main() {
	a := 3
	util.Lg.Debugf("a(binary): %08b", a)
	util.Lg.Debugf("a<<1: '%08b'", a << 1)
	util.Lg.Debugf("a<<2: '%08b'", a << 2)
	util.Lg.Debugf("a<<3: '%016b'", a << 8)
}





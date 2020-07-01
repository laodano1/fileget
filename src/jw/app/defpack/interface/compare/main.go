package main

import (
	"fileget/util"
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	z := y
	fmt.Println(z == y) //A
	util.Lg.Debugf("addr x: %p, addr y: %p", x, y)
	util.Lg.Debugf("addr y: %p, addr z: %p", y, z)
	fmt.Println(reflect.DeepEqual(x, y)) //A
}

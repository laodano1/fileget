package main

import "fileget/util"

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}


// nil is a predeclared identifier representing the zero value for a
// pointer, channel, func, interface, map, or slice type.
//var nil Type   // Type must be a pointer, channel, func, interface, map, or slice type

func main() {
	//println(InitType() == nil)
	println(InitPointer() == nil)
	println(InitEfaceType() == nil)
	t := InitEfaceType().(S)
	util.Lg.Debugf("if: %p", &t)

	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}
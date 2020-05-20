package main

import "fileget/util"

type myStatck struct {
	stack  []interface{}
	top    int32
	//Bottom int32
}

func (ms *myStatck) push(item interface{}) {
	ms.stack = append(ms.stack, item)
	ms.top++
}

func (ms *myStatck) pop() (v interface{}) {
	if ms.top < 0 {
		util.Lg.Debugf("stack empty")
		return
	}
	v = ms.stack[ms.top]
	ms.top--
	return
}

func (ms *myStatck) stackLen() (lth int32) {
	lth = ms.top + 1
	return
}

func (ms *myStatck) ShowStack() (v []interface{}) {
	v = ms.stack[:ms.top+1]
	return
}

func NewStack() (ms *myStatck) {
	ms = new(myStatck)
	ms.stack = make([]interface{}, 0)
	ms.top = -1
	return
}



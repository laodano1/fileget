package main

import "fileget/util"

func main() {
	defer func() {
		if err := recover(); err != nil {
			util.Lg.Debugf("%v", err)
		}
	}()
	ms := NewStack()
	ms.push('a')
	ms.push('A')
	ms.push("r")
	ms.push('1')
	ms.push('0')
	ms.push('9')
	util.Lg.Debugf("stack: %v", ms.ShowStack())
	util.Lg.Debugf("stack len: %v", ms.stackLen())
	//ms.pop()
	util.Lg.Debugf("pop: %v", ms.pop())
	//ms.pop()
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("pop: %v", ms.pop())
	util.Lg.Debugf("stack: %v", ms.ShowStack())
}

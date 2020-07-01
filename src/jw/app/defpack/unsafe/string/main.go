package main

import (
	"fmt"
	"runtime"
)

func main() {
	var b = []byte("123")
	//s := (*string)(unsafe.Pointer(&b))
	s := b

	b[1] = '4'
	fmt.Printf("%s, addr s: %p, addr b: %p", string(s), s, b) //print 143

	runtime.Gosched()
}

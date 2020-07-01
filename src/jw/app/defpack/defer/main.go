package main

import (
	"fmt"
)

func main() {
	fmt.Println("return:", defer_call())
}

//有名返回值则是在函数声明的同时就已经被声明，匿名返回值是在return执行时被声明，
//所以在defer语句中只能访问有名返回值，而不能直接访问匿名返回值。

//func defer_call() (ii int) {  // 有名返回值
func defer_call() int { // 匿名返回值
	var i int
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	return i
}

package main

import (
	"errors"
	"fmt"
)

func example() {
	s1 := [4]int{}
	s2 := [4]*int{}

	//s1 := make([]int, 10)
	//s2 := make([]*int, 10)
	//s3 := make([]interface{}, 10)
	a := 123
	//var a int
	s1[1] = a 	// case1: 导致a分配在栈在
	s2[1] = &a 	// case2: 导致a分配到堆上
	//s3[1] = a	// case3: 导致a分配在堆上
	//s3[1] = &a	// case4: 导致a分配在堆上
}

func ex2(i *int) *int {
	//a := i
	return i
}

func f3() {
	var err error

	defer func(e error) {
		fmt.Println(e)
	}(err)

	err = errors.New("defer error")
	return
}

func main() {
	//example()
	//i := 0
	//b := ex2(&i)
	//util.Lg.Debugf("%v", b)
	f3()


}


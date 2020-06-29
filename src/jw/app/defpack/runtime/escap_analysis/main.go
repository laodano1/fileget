package main

import (
	"errors"
	"fileget/util"
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

func t1()  {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i, val := range ss {
		ss[i] += 10
		util.Lg.Debugf("i: %v, val: %v", i, val)
	}
	for i := range ss {
		ss[i] += 10
		util.Lg.Debugf("i: %v", i)
	}

	fmt.Println(s)
}

func f1() (int, int) {
	return 1, 0
}

func t2() {
	var x int
	//x.Error()
	x = 1
	//x, _ := f1()
	util.Lg.Debugf("%v, %v", x)
	//x := errors.New("")

}

func main() {
	//example()
	//i := 0
	//b := ex2(&i)
	//util.Lg.Debugf("%v", b)
	//f3()

	t2()

}


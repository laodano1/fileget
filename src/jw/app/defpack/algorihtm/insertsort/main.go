package main

import "fmt"

func insertsort(list []int) []int {
	var cnt int
	for i := range list {
		//fmt.Printf("i: %v, val: %v\n", i, v)
		preIdx  := i - 1
		current := list[i]
		//for preIdx >= 0 && list[preIdx] > current { // asc
		for preIdx >= 0 && list[preIdx] < current {   // desc
			list[preIdx + 1] = list[preIdx]
			preIdx--
			cnt++
		}
		list[preIdx + 1] = current
	}
	fmt.Printf("cnt: %v\n", cnt)
	return list
}

func main() {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	list = insertsort(list)
	fmt.Printf("insert sort: %v\n", list)
}

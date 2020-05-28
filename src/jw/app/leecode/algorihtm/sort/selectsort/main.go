package main

import (
	"fileget/util"
)

func selectsortDesc(list []int) []int {
	for i := range list {
		max := i
		for j := i; j < len(list); j++ {
			if list[max] < list[j] {
				max, j = j, max
			}
		}
		list[i], list[max] = list[max], list[i]
	}

	return list
}

func selectsortAsc(list []int) []int {
	for i := range list {
		min := i
		for j := i; j < len(list); j++ {
			if list[min] > list[j] {
				min, j = j, min
			}
		}
		list[i], list[min] = list[min], list[i]
	}

	return list
}

func main() {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//list = selectsortDesc(list)
	list = selectsortAsc(list)
	util.Lg.Debugf("selectsort list: %v\n", list)
}

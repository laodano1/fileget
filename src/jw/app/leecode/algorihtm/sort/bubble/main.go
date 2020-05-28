package main

import (
	"fileget/util"
)

func bubble(list []int) []int {

	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
			util.Lg.Debugf("i: %v, %v", i, list)
		}
		//util.Lg.Debugf("i: %v, %v", i, list)
		util.Lg.Debugf("  ")
	}
	return list
}

func bubble2(input []int) (output []int)  {
	//for i, _ := range input {
	for i := 0; i < len(input) - 1; i++ {
		for j := 1; j < len(input) - i; j++ {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			util.Lg.Debugf("j: %v, %v", j, input)
		}
		util.Lg.Debugf("  ")
	}
	output = input
	return
}

func main() {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//list := []int{6, 0, 7, 9, 5, 4, 3, 2, 1, 8}
	//list = bubble(list)
	list = bubble2(list)
	util.Lg.Debugf("list: %v\n", list)
}

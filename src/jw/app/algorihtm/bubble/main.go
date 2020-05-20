package main

import "fmt"

func bubble(list []int) []int {

	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}


	return list
}


func main() {
	//list := []int{9, 0, 7, 6, 5, 4, 3, 2, 1, 8}
	list := []int{6, 0, 7, 9, 5, 4, 3, 2, 1, 8}
	list = bubble(list)
	fmt.Printf("list: %v\n", list)
}

package main

import "fmt"

func selectsort(list []int) []int {
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


func main() {
	list := []int{9, 0, 7, 6, 5, 4, 3, 2, 1, 8}
	list = selectsort(list)
	fmt.Printf("selectsort list: %v\n", list)
}

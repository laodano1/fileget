package main

import (
	"fileget/util"
	"math/rand"
)

func fastsort(list []int) []int {


	return list
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}


func main() {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	list = quicksort(list)
	util.Lg.Debugf("output: %v\n", list)

	a, b, c := 1, 2, 3
	util.Lg.Debugf("%v, %v, %v", a, b, c)
}

package main

import (
	"fileget/util"
)

// trade space for time
func Fastsort(input []int) (output []int) {
	//l := 0; 	r := len(input) - 1
	if len(input) <= 1 { return input}

	tmpL := make([]int, 0, len(input))
	tmpR := make([]int, 0, len(input))
	//pivot := rand.Int() % len(input)
	pivot := len(input)/2
	//pivot := 0

	for i, _ := range input {
		if i == pivot {
			continue
		} else if input[i] <= input[pivot] {  // less than and equal
			tmpL = append(tmpL, input[i])
		} else {   // greater than
			tmpR = append(tmpR, input[i])
		}
	}
	//util.Lg.Debugf("tmpL: %v", tmpL)
	//util.Lg.Debugf("tmpR: %v", tmpR)
	//util.Lg.Debugf("")

	output = append(output, Fastsort(tmpL)...)
	output = append(output, input[pivot])
	output = append(output, Fastsort(tmpR)...)

	return
}

//
func Quicksort1(a []int) []int {
	if len(a) < 2 { return a }

	left, right := 0, len(a)-1

	//pivot := rand.Int() % len(a)
	pivot := len(a)/2
	//util.Lg.Debugf("1. input: %v, pivot: %v", a, pivot)
	a[pivot], a[right] = a[right], a[pivot]
	//util.Lg.Debugf("2. input: %v", a)
	for i, _ := range a {
		if i == right {continue}
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	//util.Lg.Debugf("3. input: %v, left: %v", a, left)
	a[left], a[right] = a[right], a[left]
	//util.Lg.Debugf("4. input: %v", a)
	//util.Lg.Debugf("")
	Quicksort1(a[:left])
	Quicksort1(a[left+1:])

	return a
}


func main() {
	list := []int{11, 88, 77, 66, 9, 34, 7, 6, 5, 24, 3, 2, 1, 0}
	//list = Quicksort1(list)
	list = Fastsort(list)
	util.Lg.Debugf("output: %v\n", list)

	//a, b, c := 1, 2, 3
	//util.Lg.Debugf("%v, %v, %v", a, b, c)
}

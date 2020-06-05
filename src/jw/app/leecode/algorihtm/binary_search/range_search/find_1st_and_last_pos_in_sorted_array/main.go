package main

import (
	"fileget/util"
	"math"
)

func find_1st_and_last_pos_in_sorted_array(input []int, target int) (rg []int) {
	first := findFirst(input, target)
	last  := findFirst(input, target + 1) - 1
	//util.Lg.Debugf("last: %v", findFirst(input, target + 1) )
	if first == len(input) || input[first] != target {
		return []int{-1, -1}
	} else {
		return []int{first, int(math.Max(float64(first), float64(last)))}
	}

	return
}

func findFirst(nums []int, target int) (loc int) {
	l := 0; h := len(nums)
	for l < h {
		mid := l + (h - l)/2
		//util.Lg.Debugf("mid: %v", mid)
		if nums[mid] >= target {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func main() {
	//input := []int{3,3,3,3,4,5,5,6,7,7,7,7,8,8,10,11,11,11}
	input := []int{5,7,7,8,8,10,11}
	util.Lg.Debugf("output: %v", findFirst(input, 8))
	util.Lg.Debugf("output: %v", find_1st_and_last_pos_in_sorted_array(input, 8))
}





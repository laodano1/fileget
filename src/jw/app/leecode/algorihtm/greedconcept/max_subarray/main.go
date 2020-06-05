package main

import (
	"fileget/util"
	"math"
)

func max_subarray(input []int) int {
	if input == nil || len(input) == 0 {return 0}

	preSum := input[0]
	maxSum := preSum
	for i := 1; i < len(input); i++ {
		if preSum > 0 {
			preSum += input[i]
		} else {
			preSum = input[i]
		}
		maxSum = int(math.Max(float64(maxSum), float64(preSum)))
		util.Lg.Debugf("i: %v, masSum: %v", i, maxSum)
	}

	return maxSum
}

func main() {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	util.Lg.Debugf("output: %v", max_subarray(input))
}

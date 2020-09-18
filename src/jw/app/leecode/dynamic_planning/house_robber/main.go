package main

import "fileget/util"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func houseRobber_normal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]

}

func houseRobber_rollingArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := nums[1]

	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}

	return second
}

func houseRobber_2() { // 房子是个圈，收尾相接

}

func main() {
	input := []int{10, 3, 2, 100, 9}

	util.Lg.Debugf("max: %v", houseRobber_normal(input))

	util.Lg.Debugf("max: %v", houseRobber_rollingArray(input))

}

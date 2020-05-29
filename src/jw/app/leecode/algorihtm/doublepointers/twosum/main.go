package main

import "fileget/util"

func twosum(nums []int, target int) (idxs []int) {
	if len(nums) <= 0 {
		return
	}

	i := 0
	j := len(nums) -1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return
}


func main() {
	ipt := []int{2, 7, 11, 15}
	tgt := 18

	util.Lg.Debugf("two sum index: %v", twosum(ipt, tgt))

}

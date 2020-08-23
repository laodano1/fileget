package main

import "fileget/util"

func parent(root int) (p int) {
	//if root/2 == 0 {
	//	p = 1
	//} else {
	p = root / 2
	//}

	return
}

func left(root int) int {
	return 2 * root
}

func right(root int) int {
	return 2*root + 1
}

func setNode(nums, bhp []int, idx, root int) {
	if bhp[root] < nums[idx] { // right
		if bhp[right(root)] == 0 {
			bhp[right(root)] = nums[idx]
		} else {
			setNode(nums, bhp, idx, right(root))
		}
	} else { // left
		if bhp[left(root)] == 0 {
			bhp[left(root)] = nums[idx]
		} else {
			setNode(nums, bhp, idx, left(root))
		}
	}
}

func CreateBHeap(nums []int) []int {
	bhp := make([]int, 100)

	for i := range nums {
		if i == 0 {
			bhp[1] = nums[i]
			continue
		}
		setNode(nums, bhp, i, 0)
	}

	return bhp
}

func main() {
	//initArr := []int{5, 1, 3, 2, 6, 4}
	//initArr := []int{6, 5, 4, 3, 2, 1}
	//initArr := []int{1, 2, 3, 4, 5, 6}
	initArr := []int{3, 5, 2, 1, 4, 6}
	hp := CreateBHeap(initArr)
	util.Lg.Debugf("%v", hp)

}

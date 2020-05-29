package main

import "fileget/util"

// m, n分别为两个数组的长度
// 需要从尾开始遍历，否则在 nums1 上归并得到的值会覆盖还未进行归并比较的值。
func merge_sorted_arrays(nums1, nums2 []int, m, n int) (output []int) {
	i := m - 1; j := n -1
	k := m + n - 1

	for i >= 0 || j >= 0 {
		if i < 0 {
			nums1[k] = nums2[j]
			k--; j--
		} else if j < 0 {
			nums1[k] = nums1[i]
			k--; i--
		} else if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--; i--
		} else {
			nums1[k] = nums2[j]
			k--; j--
		}
	}

	return nums1
}


func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	util.Lg.Debugf("output: %v", merge_sorted_arrays(a, b, 3, 3))
}

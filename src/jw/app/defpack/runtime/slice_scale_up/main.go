package main

import (
	"fileget/util"
	"strings"
)

// each time, slice scales 2 times of current capacity
func main() {
	ss := []int{1,2,3,4,5}
	ret := ss[:0]
	for i := 0; i < 5; i++ {
		ret = append(ret, i)
	}
	util.Lg.Debugf("1, len: %v, cap: %v, addr ss: %p, addr ret: %p", len(ret), cap(ret), &ss, &ret)
	ret = append(ret, 6)
	util.Lg.Debugf("2, len: %v, cap: %v, addr ss: %p, addr ret: %p", len(ret), cap(ret), &ss, &ret)

	mp := make([]int, 0, 5)
	util.Lg.Debugf("3, len: %v, cap: %v", len(mp), cap(mp))
	for i := 0; i < 5; i++ {
		mp = append(mp, i)
	}
	util.Lg.Debugf("4, len: %v, cap: %v", len(mp), cap(mp))
	mp = append(mp, 6)
	util.Lg.Debugf("5, len: %v, cap: %v", len(mp), cap(mp))
	for i := 0; i < 5; i++ {
		mp = append(mp, i)
	}
	util.Lg.Debugf("6, len: %v, cap: %v", len(mp), cap(mp))
	for i := 0; i < 11; i++ {
		mp = append(mp, i)
	}
	util.Lg.Debugf("7, len: %v, cap: %v", len(mp), cap(mp))

	arr1 := [4]int{0,0,0,0}
	util.Lg.Debugf("array: %v", arr1[:0])

	var sb strings.Builder
	sb.String()

}

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fileget/util"
	"fmt"
	"io"
)

/*
	全排列
*/

var result [][]int32

func backtrack(wt []int32, w int32, track []int32) {

	canHold := func(wts, w int32) (can bool) {
		if wts > w {
			can = false
			return
		}
		return true
	}

	// if match end condition, then record the permutation
	if !canHold(wt[0], w) {
		result = append(result, track)
		return
	}

	DoExist := func(n int32) bool {
		for j := range track {
			if n == track[j] {
				return true
			}
		}
		return false
	}

	for i := range wt {
		// if path chose, continue
		if DoExist(wt[i]) {
			continue
		}

		// make choice
		track = append(track, int32(i))
		util.Lg.Debugf("1. track: %v", track)

		// next layer backtrack
		backtrack(wt, w-wt[i], track)

		track = track[:len(track)-1]
		util.Lg.Debugf("2. track: %v", track)
	}
}

// TODO: need to get maximal value
func Permutation(wt, val []int32, w int32) int32 {
	// path record
	track := make([]int32, 0)

	backtrack(wt, w, track)
	vals := make([]int32, 0)
	getTotalVal := func(vs []int32) (sum int32) {
		for i := range vs {
			sum += vs[i]
		}
		return
	}
	max := int32(0)
	for i := range result {
		if vals[i] = getTotalVal(result[i]); vals[i] > max {
			max = vals[i]
		}
	}
	return max
}

func main1() {
	wt := []int32{2, 1, 3}
	val := []int32{4, 2, 3}
	W := int32(4)

	Permutation(wt, val, W)

	for i := range result {
		util.Lg.Debugf("%v", result[i])
	}
}

func twoFunc(x int) (func(), func()) {
	return func() {
			fmt.Printf("a: %v\n", x)
		}, func() {
			fmt.Printf("b: %v\n", x)
		}
}

func main() {
	a, b := twoFunc(100)
	a()
	b()
}

func generateChallengeKey() (string, error) {
	p := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, p); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(p), nil
}

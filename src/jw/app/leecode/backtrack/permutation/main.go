package main

import "fileget/util"

/*
	全排列
*/

var result [][]int32

func backtrack(input []int32, track []int32) {
	// if match end condition, then record the permutation
	if len(track) == len(input) {
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

	for i := range input {
		// if path chose, continue
		if DoExist(input[i]) {
			continue
		}

		// make choice
		track = append(track, input[i])
		util.Lg.Debugf("1. track: %v", track)

		// next layer backtrack
		backtrack(input, track)

		track = track[:len(track)-1]
		util.Lg.Debugf("2. track: %v", track)
	}

}

func Permutation(input []int32) {
	// path record
	track := make([]int32, 0)

	backtrack(input, track)
}

func main() {
	input := []int32{1, 2, 3}
	Permutation(input)
	for i := range result {
		util.Lg.Debugf("%v", result[i])
	}
}

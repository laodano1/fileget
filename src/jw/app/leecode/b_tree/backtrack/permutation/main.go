package main

import "fileget/util"

var result [][]int32

func backtrack(input []int32, track []int32) {
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
		if DoExist(input[i]) {continue}

		// make choice
		track = append(track, input[i])

		// next layer backtrack
		backtrack(input, track)

		track = track[:len(track)-1]
	}

}

func Permutation(input []int32)  {
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

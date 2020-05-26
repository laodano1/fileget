package main

import "fileget/util"

func distribut_cookies(grid, size []int) int {
	if len(grid) <= 0 || len(size) <= 0 {return 0}

	i := 0; j := 0
	for i < len(grid) && j < len(size) {
		if grid[i] <= size[j] {
			i++
		}
		j++
	}

	return i
}

func main() {
	grid := []int{1, 3}
	size := []int{1, 2, 4}
	util.Lg.Debugf("output: %v", distribut_cookies(grid, size))
}

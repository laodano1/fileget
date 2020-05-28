package main

import (
	"fileget/util"
	"math"
)

func sumofsquarenums(ipt int) ( b bool) {
	if ipt < 0 {return false}

	i := 0
	j := math.Sqrt(float64(ipt))
	for i < int(j) {
		sum := i * i + int(j) * int(j)
		if sum == ipt {
			util.Lg.Debugf("i: %v, j: %v", i, j)
			return true
		} else if sum > ipt {
			j--
		} else {
			i++
		}

	}
	return
}

func main() {
	util.Lg.Debugf("result: %v", sumofsquarenums(300))
}

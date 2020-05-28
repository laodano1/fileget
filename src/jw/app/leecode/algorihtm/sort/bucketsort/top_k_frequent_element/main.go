package main

import "fileget/util"

func top_k_frequent_element(input []int, k int) (r []int) {
	// get each element frequency
	fre := make(map[int]int)
	var tp int
	for _, v := range input {
		fre[v] += 1
		tp++
	}
	util.Lg.Debugf("fre: %v", fre)

	// ----------- push them to buckets
	// initiate buckets
	//type bucket struct {
	//	items []int
	//}
	//buckets := make([]bucket, 0)
	buckets := make([][]int, 0)
	for i := 0; i < tp; i++ {
		bItem := make([]int, 0)
		buckets = append(buckets, bItem)
	}
	for v, cnt := range fre {
		buckets[cnt] = append(buckets[cnt], v)
	}

	// find out the top k from bucket
	t := k
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			r = append(r, buckets[i]...)
			t--
			if t <= 0 {
				break
			}
		}
	}

	return
}

func main() {
	input := []int{1,1,1,2,2,3,0,0,0,0}

	util.Lg.Debugf("output: %v", top_k_frequent_element(input, 2))
}

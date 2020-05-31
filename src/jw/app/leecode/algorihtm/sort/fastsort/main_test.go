package main

import "testing"


//BenchmarkQuicksort1-12          19724935                60.1 ns/op
//PASS
func BenchmarkQuicksort1(b *testing.B) {

	b.ResetTimer()
	list := []int{11, 88, 77, 66, 9, 34, 7, 6, 5, 24, 3, 2, 1, 0}
	for i:=0; i < b.N; i++ {
		Quicksort1(list)
	}
}

//BenchmarkFastsort-12             1000000              1048 ns/op
//PASS
func BenchmarkFastsort(b *testing.B) {

	b.ResetTimer()
	list := []int{11, 88, 77, 66, 9, 34, 7, 6, 5, 24, 3, 2, 1, 0}
	for i:=0; i < b.N; i++ {
		Fastsort(list)
	}
}
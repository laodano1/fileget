package main

import (
	"testing"
)

func BenchmarkPreorder(b *testing.B) {
	nums := []int32{2, 5, 4, 4, 1, 3}
	mt := geneTree(nums)

	for i := 0; i < b.N; i++ {
		Preorder(mt)
		//util.Lg.Debugf("%v", arr)
	}

}

func BenchmarkInorder(b *testing.B) {
	nums := []int32{2, 5, 4, 4, 1, 3}
	mt := geneTree(nums)

	for i := 0; i < b.N; i++ {
		Inorder(mt)
		//util.Lg.Debugf("%v", arr)
	}
}

func BenchmarkPostorder(b *testing.B) {
	nums := []int32{2, 5, 4, 4, 1, 3}
	mt := geneTree(nums)

	for i := 0; i < b.N; i++ {
		Postorder(mt)
		//util.Lg.Debugf("%v", arr)
	}
}

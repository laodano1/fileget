package main

import "testing"

func BenchmarkReversestring1(b *testing.B) {
	input := "aaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbccc"
	for i := 0; i < b.N; i++ {
		Reversestring1(input)
	}
}


func BenchmarkReversestring2(b *testing.B) {
	input := "aaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbcccaaabbccc"
	for i := 0; i < b.N; i++ {
		Reversestring2(input)
	}
}
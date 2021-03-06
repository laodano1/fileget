package main

import "testing"

func BenchmarkGetUnit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUnit()
	}
}

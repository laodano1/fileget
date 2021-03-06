package main

import "testing"

func BenchmarkCallRPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallRPC()
	}
}

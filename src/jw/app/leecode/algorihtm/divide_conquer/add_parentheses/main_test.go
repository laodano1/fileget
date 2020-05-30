package main

import "testing"

func BenchmarkValueOfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValueOfString("222- 122+ 133")
	}
}



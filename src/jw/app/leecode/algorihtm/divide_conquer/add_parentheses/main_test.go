package main

import (
	"fileget/util"
	"testing"
)

func BenchmarkValueOfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.ValueOfString("222- 122+ 133")
	}
}



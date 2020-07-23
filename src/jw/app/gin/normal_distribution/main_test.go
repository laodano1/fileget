package main

import (
	"fileget/util"
	"testing"
)

func TestGetNormFloat64(t *testing.T) {

}

func BenchmarkGetNormFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.Lg.Debugf("rand: %v", GetNormFloat64(1000, 15000))
	}
}

func BenchmarkGetVal(b *testing.B) {
	tR := &tmpRands{}
	for i := 0; i < b.N; i++ {
		GenerateVal(tR)
		mid := int64(len(tR.Arr)/2)
		if tR.Arr[mid] < 1000 || tR.Arr[mid] > 2000 {
			util.Lg.Debugf("********************** bad value!!!")
		}
		//util.Lg.Debugf("rand: %v", tR.Arr[int64(len(tR.Arr)/2)])
	}
}

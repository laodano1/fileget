package main

import (
	"fileget/util"
	"math"
	"testing"
)

func TestGetNormFloat64(t *testing.T) {
	util.Lg.Debugf("%v", math.Ceil(8.1))   // 向上取整
	util.Lg.Debugf("%v", math.Round(8.51)) // 四舍五入

	tmp := 123456

	util.Lg.Debugf("%v", tmp - tmp % 10000) //

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

func BenchmarkGetVal2(b *testing.B) {
	min := float64(2500000)
	max := float64(3000000)
	for i := 0; i < b.N; i++ {
		util.Lg.Debugf("value: %v", GenerateVal2(min, max))
	}
}
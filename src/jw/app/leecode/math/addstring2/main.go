package main

import (
	"fileget/util"
)

func ParseStr2Int(i string) int {
	chars := []byte(i)
	sum := 0
	bt := len(chars) - 1
	for _, c := range chars {
		tmp := int(c - '0')
		util.Lg.Debugf("tmp: %d", tmp)
		//if bt <= 1 {
		//	sum += tmp
		//} else {
		//	sum += tmp * bt * 10
		//
		//
		//}
		bt--
	}
	return sum
}

func main() {
	str1 := "123"
	util.Lg.Debugf("sum: %v", ParseStr2Int(str1))
}

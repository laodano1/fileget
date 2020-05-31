package main

import "fileget/util"

func Reversestring1(input string) (s string) {
	rs := []rune(input)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	s = string(rs)
	return
}

func Reversestring2(input string) (s string) {
	for _, v := range input {
		s = string(v) + s
	}

	return
}

func main() {
	//input := "aaabbccc"
	//util.Lg.Debugf("output: %v", Reversestring1(input))
	//util.Lg.Debugf("output: %v", Reversestring2(input))
	util.Lg.Debugf("A: %d", 'A')
	util.Lg.Debugf("a: %d", 'a')

}

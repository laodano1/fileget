package main

import (
	"fileget/util"
	"strings"
)

var (
	vowles = map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
)

func reversvowalofastring(input string) (output string) {
	if input == "" {return ""}
	tmp := make([]string, len(input))

	i := 0
	j := len(input) - 1
	for i <= j {
		ci := input[i:i+1]
		cj := input[j:j+1]
		if _, ok := vowles[ci]; !ok {
			tmp[i] = ci; i++
		} else if _, ok := vowles[cj]; !ok {
			tmp[j] = cj; j--
		} else {
			tmp[i] = cj; i++
			tmp[j] = ci; j--
		}
	}

	output = strings.Join(tmp, "")
	return
}

func main() {

	util.Lg.Debugf("output: %v", reversvowalofastring("leetcode"))
	util.Lg.Debugf("output: %v", reversvowalofastring("hello"))
}

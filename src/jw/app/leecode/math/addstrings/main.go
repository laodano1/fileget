package main

import (
	"fileget/util"
	"fmt"
	"strings"
)

func addstrings(num1, num2 string) ( s string) {
	tmp    := make([]string, 0)
	result := make([]string, 0)
	l1 := len(num1) - 1; 	l2 := len(num2) - 1
	c := 0
	rc := rune(c)

	b1 := []rune(num1)
	b2 := []rune(num2)

	//util.Lg.Debugf("a: %v", string(5))
	for rc > 0 || l1 >= 0 || l2 >= 0 {
		var x rune
		var y rune
		if l1 < 0 { x = 0} else { x = (b1[l1]) - '0' }
		if l2 < 0 { y = 0} else { y = (b2[l2]) - '0' }
		l1--; l2--
		tmp = append(tmp, fmt.Sprintf("%d", (x + y + rc) % 10))
		rc = (x + y + rc) / 10
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		result = append(result, tmp[i])
	}

	s = strings.Join(result, "")
	return
}

func main() {
	util.Lg.Debugf("output: %d", '0')
	util.Lg.Debugf("output: %d", '0' - '0')
	util.Lg.Debugf("output: %d", '9' - '0')

	n1 :=  "999"
	n2 := "9999"

	util.Lg.Debugf("output: %v", addstrings(n1, n2))

}

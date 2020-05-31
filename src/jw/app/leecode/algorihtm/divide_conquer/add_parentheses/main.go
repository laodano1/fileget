package main

import (
	"fileget/util"
)

func add_parentheses(input string) []int {
	ways := make([]int, 0)
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if c == '+' || c == '-' || c == '*' {
			left  := add_parentheses(input[:i])
			right := add_parentheses(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch c {
					case '+':
						ways = append(ways, l + r)
					case '-':
						ways = append(ways, l - r)
					case '*':
						ways = append(ways, l * r)
					}
				}
			}
		}
	}
	if len(ways) == 0 {
		ways = append(ways, util.ValueOfString(input))
	}

	return ways
}


func main() {
	//util.Lg.Debugf("value: %v", ValueOfString("222-122+133"))
	//util.Lg.Debugf("value: %v", ValueOfString(" 3 * 2 + 1"))
	util.Lg.Debugf("value: %v", add_parentheses("2-1-1"))
	util.Lg.Debugf("value: %v", add_parentheses("2*3-4*5"))
}

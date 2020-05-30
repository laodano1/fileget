package main

import (
	"fileget/util"
	"strconv"
	"strings"
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
		ways = append(ways, ValueOfString(input))
	}

	return ways
}

func ValueOfString(input string) int {
	if len(input) == 0 {return 0}

	rs := []rune(input)
	sb := make([]rune, 0)
	nbs := make([][]string, 0)
	for i, j := 0, 0; i < len(input); i++ {
		switch rs[i] {
		case ' ':
			continue
		case '+', '-', '*':
			sb = append(sb, rs[i])
		default:
			if i == 0 {
				tmp := make([]string, 0)
				tmp = append(tmp, string(rs[i]))
				nbs = append(nbs, tmp)
			} else {
				if rs[i-1] == '+' || rs[i-1] == '-' ||rs[i-1] == '*' || rs[i-1] == ' ' {
					tmp := make([]string, 0)
					tmp = append(tmp, string(rs[i]))
					nbs = append(nbs, tmp)
					j++
				}  else {
					nbs[j] = append(nbs[j], string(rs[i]))
				}
			}
		}
	}
	//util.Lg.Debugf("sb: %c", sb)
	//util.Lg.Debugf("nbs: %v", nbs)

	var sum int
	nums := make([]int, 0)
	for i := range nbs {
		t1, _ := strconv.Atoi(strings.Join(nbs[i], ""))
		nums = append(nums, t1)
	}

	for i, j := 0, 0; i < len(nums); i++ {
		if i == 0 {
			sum = nums[i]; continue
		} else {
			sum = calculate(sum, nums[i], sb[j])
			//util.Lg.Debugf("sum: %v", sum)
			j++
		}
	}

	return sum
}

func calculate(a, b int, s rune) (sum int) {
	if s == '+' {
		sum = a + b
	} else if s == '-' {
		sum = a - b
	} else {
		sum = a * b
	}
	return
}

func main() {
	//util.Lg.Debugf("value: %v", ValueOfString("222-122+133"))
	//util.Lg.Debugf("value: %v", ValueOfString(" 3 * 2 + 1"))
	util.Lg.Debugf("value: %v", add_parentheses("2-1-1"))
	util.Lg.Debugf("value: %v", add_parentheses("2*3-4*5"))
}

package main

import "fileget/util"

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--; r++
	}
	util.Lg.Debugf("l: %v, r: %v", l, r)
	tmp := s[l+1:r-l-1]
	util.Lg.Debugf("'%v'", tmp)
	return tmp
}

func longestPalindrome(s string) (res string) {
	if s == "" {return s}

	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i + 1)
		if len(res) > len(s1) { res = res } else { res = s1 }
		if len(res) > len(s2) { res = res } else { res = s2 }
	}
	return
}

func main() {
	s := "abbaycaa"
	longestPalindrome(s)
}
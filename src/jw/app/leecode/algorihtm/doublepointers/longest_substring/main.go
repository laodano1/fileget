package main

import "fileget/util"

func isSubString(s, target string) bool {
	i := 0; j := 0
	for i < len(s) && j < len(target) {
		if s[i:i+1] == target[j:j+1] {
			j++
		}
		i++
	}

	return j == len(target)
}

func longest_substring(s string, d []string) string {
	longestWord := ""
	for _, target := range d {
		l1 := len(longestWord); l2 := len(target)
		if l1 > l2 || (l1 == l2 && longestWord != target) {
			continue
		} 
		if isSubString(s, target) {
			longestWord = target
		}
	}
	return longestWord
}

func main() {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}
	util.Lg.Debugf("output: %v", longest_substring(s, d))
}

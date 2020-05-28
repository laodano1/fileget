package main

import "fileget/util"

func isPalindrome(input string, i, j int) bool {
	for i < j {
		if input[i:i+1] != input[j:j+1] {
			return false
		}
		i++; j--
	}

	return true
}

func verify_huiwen_palindrome(input string) bool {
	i := 0
	j := len(input) - 1

	for i < j {
		if input[i:i+1] != input[j:j+1] {
			return isPalindrome(input, i, j-1) || isPalindrome(input, i+1, j)
		}
		i++; j--
	}

	return true
}


func main() {

	util.Lg.Debugf("is palindrom: %v", verify_huiwen_palindrome("abca"))
	util.Lg.Debugf("is palindrom: %v", verify_huiwen_palindrome("abacda"))

}

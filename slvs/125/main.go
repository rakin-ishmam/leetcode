package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/description/

var str string

func isal(b byte) bool {
	if b < 'a' || b > 'z' {
		return false
	}

	return true
}

func isnum(b byte) bool {
	if b < '0' || b > '9' {
		return false
	}

	return true
}

func isPal(i, j int) bool {
	if i >= j {
		return true
	}

	if !isal(str[i]) && !isnum(str[i]) {
		return isPal(i+1, j)
	}

	if !isal(str[j]) && !isnum(str[j]) {
		return isPal(i, j-1)
	}

	if str[i] != str[j] {
		return false
	}

	return isPal(i+1, j-1)
}

func isPalindrome(s string) bool {
	str = strings.ToLower(s)
	return isPal(0, len(str)-1)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

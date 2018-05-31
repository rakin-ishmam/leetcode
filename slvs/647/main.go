package main

// https://leetcode.com/problems/palindromic-substrings/description/

import "fmt"

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	cnt := 0

	for i := range s {
		cnt += palcnt(i, i, s) + palcnt(i, i+1, s)
	}

	return cnt
}

func palcnt(l, r int, str string) int {
	cnt := 0

	for ; l >= 0 && r < len(str) && str[l] == str[r]; l, r = l-1, r+1 {
		cnt++
	}

	return cnt
}

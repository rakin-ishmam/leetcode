package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring/description/

var bstr []byte
var mem [][]int

func initMem() {
	n := len(bstr)
	mem = make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}
}

func f(l, r int) int {
	if l >= r {
		return 1
	}

	if mem[l][r] > -1 {
		return mem[l][r]
	}

	if bstr[l] == bstr[r] {
		mem[l][r] = f(l+1, r-1)
		return mem[l][r]
	}

	mem[l][r] = 0
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func longestPalindrome(s string) string {
	bstr = []byte(s)
	initMem()

	res := 0
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if f(i, j) == 1 {
				diff := j - i + 1
				if diff > res {
					res = diff
					l = i
					r = j
				}

			}
		}
	}

	return s[l : r+1]
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

package main

// https://leetcode.com/problems/delete-operation-for-two-strings/description/

import (
	"fmt"
)

var mem [][]int

func initMem(n int) {
	mem = make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}
}

func f(s1, s2 string, i, j int) int {
	if i >= len(s1) || j >= len(s2) {
		return 0
	}

	if mem[i][j] > -1 {
		return mem[i][j]
	}

	res := f(s1, s2, i+1, j)

	r1 := f(s1, s2, i, j+1)
	if r1 > res {
		res = r1
	}

	if s1[i] == s2[j] {
		r2 := f(s1, s2, i+1, j+1) + 1
		if r2 > res {
			res = r2
		}
	}

	mem[i][j] = res
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func minDistance(word1 string, word2 string) int {
	initMem(max(len(word1), len(word2)))
	mx := f(word1, word2, 0, 0)
	fmt.Println(mx)
	return len(word1) - mx + len(word2) - mx
}

func main() {
	fmt.Println("yo", minDistance("sea", "eat"))
}

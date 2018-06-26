package main

// https://leetcode.com/problems/repeated-string-match/description/
import "fmt"

var a, b string

func can(i, j int, started bool) int {

	if j >= len(b) {
		return 0
	}

	if i >= len(a) {
		if !started {
			return -1
		}

		ans := can(0, j, started)
		if ans < 0 {
			return ans
		}

		return ans + 1
	}

	if a[i] == b[j] {
		return can(i+1, j+1, true)
	}

	if started {
		return -1
	}

	return can(i+1, j, false)
}

func repeatedStringMatch(A string, B string) int {
	a, b = A, B
	for i := range A {
		ans := can(i, 0, false)
		if ans >= 0 {
			return ans + 1
		}
	}

	return -1
}

func main() {

	fmt.Println("yo", repeatedStringMatch("abababaaba", "aabaaba"))
}

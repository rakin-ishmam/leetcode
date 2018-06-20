package main

// https://leetcode.com/problems/k-th-symbol-in-grammar/description/
import "fmt"

var sum []int

func initSum() {
	sum = make([]int, 30)
	sum[0] = 1
	for i := 1; i < 30; i++ {
		sum[i] = 2 * sum[i-1]
	}
}

func f(cnt, k int) int {
	// fmt.Println(k)
	if k == 1 {
		if cnt%2 == 0 {
			return 0
		}
		return 1
	}
	if k == 2 {
		if cnt%2 == 0 {
			return 1
		}
		return 0
	}

	for i := 0; i < 30; i++ {
		if k <= sum[i] {
			return f(cnt+1, k-sum[i-1])
		}
	}

	return 0
}

func kthGrammar(N int, K int) int {
	initSum()
	return f(0, K)
}

func main() {
	fmt.Println(kthGrammar(30, 3))
}

// 1, 2, 4, 8, 16
// 0
// 01
// 0110
// 01101001
// 0110100110010110
// 0 1 1 0 1 0 0 1 1 0  0  1  0  1  1  0  1  0  0  1  0  1  1  0  0  1  1  0  1  0  0  1
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32

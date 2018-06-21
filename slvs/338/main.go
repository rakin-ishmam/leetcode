package main

// https://leetcode.com/problems/counting-bits/description/
import "fmt"

func countBits(num int) []int {
	sol := make([]int, num+1)

	mod := 1
	for i := 0; i <= num; i++ {
		if i <= 1 {
			sol[i] = i
			continue
		}

		if i == mod*2 {
			mod = i
		}

		sol[i] = 1 + sol[i%mod]

	}

	return sol
}

func main() {
	fmt.Println(countBits(5))
}

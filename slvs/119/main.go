package main

// https://leetcode.com/problems/pascals-triangle-ii/description/
import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}
	aa := make([][]int, 2)
	aa[0] = make([]int, rowIndex+1)
	aa[1] = make([]int, rowIndex+1)
	aa[1][0], aa[1][1] = 1, 1

	p := 1
	c := 0

	for i := 2; i <= rowIndex; i++ {
		aa[c][0], aa[c][i] = 1, 1
		for j := 1; j < i; j++ {
			aa[c][j] = aa[p][j-1] + aa[p][j]
		}
		p, c = c, p
	}

	return aa[p]
}

func main() {
	fmt.Println("yo world", getRow(7))
}

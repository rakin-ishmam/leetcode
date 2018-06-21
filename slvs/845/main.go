package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/longest-mountain-in-array/description/

var linc, rinc []int

func genlinc(aa []int) {
	linc = make([]int, len(aa))
	for i := 1; i < len(aa); i++ {
		if aa[i] > aa[i-1] {
			linc[i] = 1 + linc[i-1]
			if linc[i-1] == 0 {
				linc[i]++
			}
		}
	}
}

func genrinc(aa []int) {
	rinc = make([]int, len(aa))
	for i := len(aa) - 2; i >= 0; i-- {
		if aa[i] > aa[i+1] {
			rinc[i] = 1 + rinc[i+1]
			if rinc[i+1] == 0 {
				rinc[i]++
			}
		}
	}
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func longestMountain(A []int) int {
	genlinc(A)
	genrinc(A)

	res := 0

	for i := range A {
		if linc[i] == 0 || rinc[i] == 0 {
			continue
		}
		res = max(res, linc[i]+rinc[i]-1)
	}

	if res == 1 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(longestMountain([]int{0, 2, 4, 6, 0}))

}

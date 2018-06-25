package main

// https://leetcode.com/problems/card-flipping-game/description/

import (
	"fmt"
	"sort"
)

func flipgame(fronts []int, backs []int) int {
	dupCnt := make([]int, 2002)
	all := []int{}
	for i := 0; i < len(fronts); i++ {
		all = append(all, fronts[i], backs[i])

		if fronts[i] == backs[i] {
			dupCnt[fronts[i]]++
		}
	}
	sort.Ints(all)
	for _, v := range all {
		if dupCnt[v] > 0 {
			continue
		}

		return v
	}

	return 0
}

func main() {
	fmt.Println("yo world", flipgame([]int{1, 1}, []int{1, 2}))
}

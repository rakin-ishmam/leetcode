package main

// https://leetcode.com/problems/partition-labels/description/
import (
	"fmt"
	"math"
)

func partitionLabels(S string) []int {
	rmost := findRMost(S)

	res := []int{}
	var r int
	cnt := 0
	for i := 0; i < len(S); i++ {
		if cnt == 0 {
			cnt = 1
			r = rmost[S[i]-'a']
		}

		if i == r {
			res = append(res, cnt)
			cnt = 0
			continue
		}

		cnt++
		r = int(math.Max(float64(r), float64(rmost[S[i]-'a'])))
	}

	if cnt > 0 {
		res = append(res, cnt)
	}
	return res
}

func findRMost(s string) []int {
	rmost := make([]int, 26)
	for i := range s {
		rmost[s[i]-'a'] = i
	}

	return rmost
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

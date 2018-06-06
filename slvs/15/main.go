package main

// https://leetcode.com/problems/3sum/description/

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			c := -a - b

			if _, ok := m[c]; ok {
				res = append(res, []int{a, b, c})

				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
			}
			if j < len(nums) {
				m[nums[j]] = true
			}
		}
	}

	return res
}

func main() {
	res := threeSum([]int{-2, 0, 1, 1, 2})
	fmt.Println(len(res), res)
}

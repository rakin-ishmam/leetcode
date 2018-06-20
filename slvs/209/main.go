package main

// https://leetcode.com/problems/minimum-size-subarray-sum/description/

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, 0
	sum := nums[0]
	if sum >= s {
		return 1
	}
	res := len(nums) + 1
	for r < len(nums) {
		if sum < s || r < l {
			r++
			if r < len(nums) {
				sum += nums[r]
			}
			continue
		}
		res = min(res, r-l+1)

		sum -= nums[l]
		l++
	}

	if res > len(nums) {
		return 0
	}
	return res
}

func main() {
	fmt.Println(minSubArrayLen(3, []int{}))
}

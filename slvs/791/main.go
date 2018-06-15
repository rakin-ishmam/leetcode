package main

// https://leetcode.com/problems/custom-sort-string/description/

import (
	"fmt"
	"sort"
)

type node struct {
	c      byte
	weight int
}

type nodes []node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].weight < n[j].weight
}

func (n *nodes) Swap(i, j int) {
	nn := *n
	nn[i], nn[j] = nn[j], nn[i]
}

func customSortString(S string, T string) string {
	w := weight(S)
	ns := nodes{}

	for i := range T {
		n := node{
			c:      T[i],
			weight: w[T[i]-'a'],
		}

		ns = append(ns, n)
	}

	sort.Sort(&ns)

	res := []byte{}
	for _, v := range ns {
		res = append(res, v.c)
	}

	return string(res)
}

func btToInt(b byte) int {
	return int(b - 'a')
}

func weight(s string) []int {
	aa := make([]int, 26)
	for i := range aa {
		aa[i] = 100
	}

	for i := range []byte(s) {
		aa[btToInt(s[i])] = i
	}

	return aa
}

func main() {
	fmt.Println(customSortString("cba", "abcd"))
}

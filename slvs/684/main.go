package main

// https://leetcode.com/problems/redundant-connection/description/

import "fmt"

var grph [][]int
var visited []bool

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func reset() {
	for i := range visited {
		visited[i] = false
	}
}

func initg(edges [][]int) {
	n := 0
	for _, e := range edges {
		n = max(n, e[0])
		n = max(n, e[1])
	}

	grph = make([][]int, n)
	visited = make([]bool, n)

	for _, e := range edges {
		addE(e[0]-1, e[1]-1)
	}

	reset()
}

func addE(e1, e2 int) {
	grph[e1] = append(grph[e1], e2)
	grph[e2] = append(grph[e2], e1)
}

func walk(i, e1, e2 int) {
	if visited[i] {
		return
	}

	visited[i] = true
	for _, v := range grph[i] {
		if (i == e1 && v == e2) || (i == e2 && v == e1) {
			continue
		}

		walk(v, e1, e2)
	}

}

func findRedundantConnection(edges [][]int) []int {
	initg(edges)

	for i := len(edges) - 1; i >= 0; i-- {
		e := edges[i]
		reset()

		walk(e[0]-1, e[0]-1, e[1]-1)
		cnt := 0
		for _, v := range visited {
			if !v {
				break
			}
			cnt++
		}

		if cnt == len(visited) {
			return e
		}
	}

	return []int{}

}

func main() {
	fmt.Println(findRedundantConnection([][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
		[]int{1, 4},
		[]int{1, 5},
	}))
}

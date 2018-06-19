package main

import "fmt"

// https://leetcode.com/problems/keys-and-rooms/description/

var visited []bool

func initVis(n int) {
	visited = make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
}

var grph [][]int

func walk(n int) {
	if visited[n] {
		return
	}

	visited[n] = true

	for i := 0; i < len(grph[n]); i++ {
		walk(grph[n][i])
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	initVis(len(rooms))
	grph = rooms
	walk(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("yo world")
}

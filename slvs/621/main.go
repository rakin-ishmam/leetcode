package main

// https://leetcode.com/problems/task-scheduler/description/
import (
	"fmt"
	"sort"
)

type item struct {
	task int
	tot  int
}

type items []item

func (is items) Len() int {
	return len(is)
}

func (is items) Less(i, j int) bool {
	return is[i].tot > is[j].tot
}

func (is *items) Swap(i, j int) {
	iss := *is
	iss[i], iss[j] = iss[j], iss[i]
}

func leastInterval(tasks []byte, n int) int {
	chs := make([]int, 26)
	for _, b := range tasks {
		chs[b-'A']++
	}

	its := items{}
	for i, v := range chs {
		if v == 0 {
			continue
		}

		its = append(its, item{task: i, tot: v})
	}

	sort.Sort(&its)
	totInt := (its[0].tot - 1) * n
	// fmt.Println(totInt, n)

	for i := 1; i < len(its); i++ {
		if its[i].tot < its[0].tot {
			totInt -= its[i].tot
		} else {
			totInt -= its[i].tot - 1
		}
	}

	// fmt.Println(totInt, n, len(tasks))

	if totInt <= 0 {
		return len(tasks)
	}

	return len(tasks) + totInt
}

func main() {
	fmt.Println(leastInterval([]byte("AAABBB"), 2))
}

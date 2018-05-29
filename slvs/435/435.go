package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/non-overlapping-intervals/description/

//  Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type intervals []Interval

func (is *intervals) Len() int {
	return len(*is)
}

func (is intervals) Less(i, j int) bool {
	if is[i].Start != is[i].Start {
		return is[i].Start < is[j].Start
	}

	return is[i].End < is[j].End
}

func (is *intervals) Swap(i, j int) {
	ar := *is
	ar[i], ar[j] = ar[j], ar[i]
}

type node struct {
	val Interval
	pre *node
}

type stack struct {
	node *node
	l    int
}

func (s *stack) push(v Interval) {
	n := node{val: v}
	if s.node != nil {
		n.pre = s.node
	}

	s.node = &n
	s.l++
}

func (s stack) len() int {
	return s.l
}

func (s *stack) pop() {
	if s.node == nil {
		return
	}

	s.node = s.node.pre
	s.l--
}

func (s *stack) top() Interval {
	if s.node == nil {
		return Interval{}
	}

	return s.node.val
}

func eraseOverlapIntervals(invs []Interval) int {
	is := intervals(invs)
	sort.Sort(&is)

	cnt := 0
	s := stack{}

	for _, inv := range is {
		if s.len() == 0 {
			s.push(inv)
			continue
		}

		sinv := s.top()
		if sinv.End <= inv.Start {
			s.push(inv)
			continue
		}

		if sinv.End > inv.End {
			s.pop()
			s.push(inv)
			cnt++
		} else {
			cnt++
		}
	}

	return cnt
}

func main() {

	invs := intervals{{1, 2}, {2, 3}}
	fmt.Println(eraseOverlapIntervals(invs))

}

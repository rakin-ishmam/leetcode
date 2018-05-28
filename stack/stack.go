package main

import "fmt"

type node struct {
	val int
	pre *node
}

type stack struct {
	node *node
	l    int
}

func (s *stack) push(v int) {
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

func (s *stack) top() int {
	if s.node == nil {
		return 0
	}

	return s.node.val
}

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.len())
	fmt.Println(s.top())
	s.pop()
	fmt.Println(s.top())
	s.pop()
	fmt.Println(s.len())
}

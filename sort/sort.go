package main

import (
	"fmt"
	"sort"
)

type person struct {
	age  int
	name string
}

type persons []person

func (p *persons) Len() int {
	return len(*p)
}

func (p persons) Less(i, j int) bool {
	return p[i].age < p[j].age
}

func (p *persons) Swap(i, j int) {
	pp := *p
	pp[i], pp[j] = pp[j], pp[i]
}

func main() {
	ps := persons{}
	ps = append(ps, person{10, "a"})
	ps = append(ps, person{4, "b"})
	ps = append(ps, person{2, "c"})
	sort.Sort(&ps)
	fmt.Println(ps)
}

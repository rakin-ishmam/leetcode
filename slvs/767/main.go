package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/reorganize-string/description/

type node struct {
	ch    int
	count int
}

type nodes []node

func (n *nodes) Len() int {
	return len(*n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].count < n[j].count
}

func (n *nodes) Swap(i, j int) {
	nn := *n
	nn[i], nn[j] = nn[j], nn[i]
}

func mknodes() nodes {
	ns := nodes(make([]node, 26))

	for i := range ns {
		ns[i].ch = i
	}

	return ns
}

func setcnt(ns *nodes, str string) {
	vns := *ns
	for i := range str {
		vns[str[i]-'a'].count++
	}
}
func uchar(ns *nodes) int {
	vns := *ns
	for i := 25; i >= 0; i-- {
		if vns[i].count == 0 {
			return 25 - i
		}
	}

	return 26
}
func reorganizeString(S string) string {
	ns := mknodes()
	setcnt(&ns, S)

	sort.Sort(&ns)
	mxcnt := ns[25].count
	if mxcnt > (len(S)+1)/2 {
		return ""
	}

	pos := make([][]byte, mxcnt)
	ucnt := uchar(&ns)

	posind := 0

	for i := 0; i < ucnt; i++ {
		for j := 0; j < ns[25-i].count; j++ {
			pos[posind] = append(pos[posind], byte(ns[25-i].ch+'a'))
			posind++
			if posind >= mxcnt {
				posind = 0
			}
		}
	}

	bts := []byte{}

	for i := 0; i < mxcnt; i++ {
		bts = append(bts, pos[i]...)
	}

	return string(bts)
}

func main() {
	fmt.Println(reorganizeString("aaabc"))
}

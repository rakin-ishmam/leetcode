package main

// https://leetcode.com/problems/longest-common-prefix/description/
import "fmt"

var ss []string

func prefix(i int) []byte {
	bb := []byte{}
	if len(ss) == 0 {
		return bb
	}

	var b byte
	for si, s := range ss {
		if i >= len(s) {
			return bb
		}

		if si == 0 {
			b = s[i]
			continue
		}

		if s[i] != b {
			return bb
		}
		break
	}

	bb = append(bb, b)
	bb = append(bb, prefix(i+1)...)

	return bb
}

func longestCommonPrefix(strs []string) string {
	ss = strs
	return string(prefix(0))

}

func main() {
	fmt.Println("yo world", longestCommonPrefix([]string{}))
}

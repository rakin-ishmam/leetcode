package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/multiply-strings/description/

func mxmnstr(s1, s2 string) (mx, mn string) {
	mx, mn = s1, s2
	if len(s2) > len(s1) {
		mx = s2
		mn = s1
	}
	return
}

func chToInt(b byte) int {
	return int(b - '0')
}

func intToCh(i int) byte {
	return byte(i) + '0'
}

func rev(bb []byte) []byte {
	nxtbb := make([]byte, len(bb))

	for i, j := 0, len(bb)-1; i < len(bb); i, j = i+1, j-1 {
		nxtbb[i] = bb[j]
	}

	return nxtbb
}

func calRem(num int) (rem, val int) {
	if num < 10 {
		val = num
		rem = 0
		return
	}

	rem = num / 10
	val = num % 10
	return
}

func multiplyStr(mx string, num int) string {
	res := []byte{}

	rem := 0
	for i := len(mx) - 1; i >= 0; i-- {
		cur := (chToInt(mx[i]) * num) + rem
		rem, cur = calRem(cur)
		// fmt.Println(mx, num, rem, cur)
		res = append(res, intToCh(cur))
	}
	if rem > 0 {
		res = append(res, intToCh(rem))
	}

	return string(rev(res))
}

func addstr(s1, s2 string) string {
	b1 := rev([]byte(s1))
	b2 := rev([]byte(s2))
	res := []byte{}

	rem := 0
	for i := 0; i < len(s1); i++ {
		val := chToInt(b1[i]) + chToInt(b2[i]) + rem
		rem, val = calRem(val)
		res = append(res, intToCh(val))
	}

	if rem > 0 {
		res = append(res, intToCh(rem))
	}

	return string(rev(res))
}

func add(up, lw string, offset int) string {
	nup := up
	nlw := fmt.Sprintf("%s%s", lw, strings.Repeat("0", offset))

	if len(nlw) < len(nup) {
		nlw = fmt.Sprintf("%s%s", strings.Repeat("0", len(nup)-len(nlw)), nlw)
	}

	if len(nlw) > len(nup) {
		nup = fmt.Sprintf("%s%s", strings.Repeat("0", len(nlw)-len(nup)), nup)
	}

	return addstr(nup, nlw)
}

func multiply(num1 string, num2 string) string {
	mx, mn := mxmnstr(num1, num2)
	res := []byte{}

	off := 0
	for mni := len(mn) - 1; mni >= 0; mni-- {
		val := multiplyStr(mx, int(mn[mni]-'0'))
		res = []byte(add(string(res), val, off))
		off++
	}

	if res[0] == '0' {
		return "0"
	}

	return string(res)

}

func main() {
	fmt.Println(multiply("0", "99999"))
}

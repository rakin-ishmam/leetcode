package main

// https://leetcode.com/problems/add-binary/description/
import (
	"fmt"
	"strings"
)

func eqlen(a, b string) (string, string) {
	aa := a
	bb := b
	if len(a) < len(b) {
		aa = fmt.Sprintf("%s%s", strings.Repeat("0", len(b)-len(a)), aa)
	}
	if len(b) < len(a) {
		bb = fmt.Sprintf("%s%s", strings.Repeat("0", len(a)-len(b)), bb)
	}

	return aa, bb
}

func add(rem, a1, a2 byte) (byte, byte) {
	sum := 0
	sum += int(a1 - '0')
	sum += int(a2 - '0')
	sum += int(rem - '0')

	switch sum {
	case 0:
		return '0', '0'
	case 1:
		return '0', '1'
	case 2:
		return '1', '0'
	default:
		return '1', '1'
	}
}

func rev(bb []byte) []byte {
	rbb := make([]byte, len(bb))

	for i, j := 0, len(bb)-1; i < len(bb); i, j = i+1, j-1 {
		rbb[i] = bb[j]
	}

	return rbb
}

func addBinary(a string, b string) string {
	a, b = eqlen(a, b)
	ans := []byte{}

	var rem, v byte
	rem = '0'
	v = '0'
	for i := len(a) - 1; i >= 0; i-- {
		rem, v = add(rem, a[i], b[i])
		ans = append(ans, v)
	}

	if rem == '1' {
		ans = append(ans, rem)
	}

	return string(rev(ans))
}

func main() {
	fmt.Println("yo world", addBinary("1010", "1011"))
}

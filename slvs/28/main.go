package main

// https://leetcode.com/problems/implement-strstr/description/
import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println("yo world", strStr("hello", ""))
}

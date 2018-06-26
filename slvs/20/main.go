package main

// https://leetcode.com/problems/valid-parentheses/description/
import "fmt"

type stack struct {
	bts []byte
}

func (s *stack) push(b byte) {
	s.bts = append(s.bts, b)
}

func (s *stack) len() int {
	return len(s.bts)
}

func (s *stack) pop() byte {
	b := s.bts[len(s.bts)-1]
	s.bts = s.bts[:len(s.bts)-1]
	return b
}

func isop(b byte) bool {
	switch b {
	case '[':
		return true
	case '{':
		return true
	case '(':
		return true
	}

	return false
}

func getop(b byte) byte {
	switch b {
	case ']':
		return '['
	case '}':
		return '{'
	case ')':
		return '('
	}

	return '3'
}

func isValid(s string) bool {
	bb := []byte(s)
	st := stack{[]byte{}}

	for _, b := range bb {
		if isop(b) {
			st.push(b)
			continue
		}

		if st.len() == 0 {
			return false
		}

		op := st.pop()
		if op != getop(b) {
			return false
		}
	}

	if st.len() > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println("yo world", isValid("[]"))
}

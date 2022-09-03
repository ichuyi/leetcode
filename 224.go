package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	ok, v := getNumber(s)
	if ok {
		return v
	}
	result := 0
	var operator rune
	for offset := 0; offset < len(s); {
		switch s[offset] {
		case '(':
			number := 1
			offset++
			start := offset
			for number != 0 {
				switch s[offset] {
				case '(':
					number++
				case ')':
					number--
				default:
				}
				offset++
			}
			n := calculate(s[start : offset-1])
			if operator == '+' {
				result = result + n
			} else if operator == '-' {
				result = result - n
			} else {
				result = n
			}
		case '-', '+':
			operator = rune(s[offset])
			offset++
		default:
			start := offset
			for ; offset < len(s) && isNumber(rune(s[offset])); offset++ {
			}
			_, n := getNumber(s[start:offset])
			if operator == '+' {
				result = result + n
				operator = 0
			} else if operator == '-' {
				result = result - n
				operator = 0
			} else {
				result = n
			}
		}
	}
	return result
}
func getNumber(i string) (bool, int) {
	v, err := strconv.Atoi(i)
	if err != nil {
		return false, 0
	}
	return true, v
}
func isNumber(i rune) bool {
	m := map[rune]struct{}{
		'(': {},
		')': {},
		'+': {},
		'-': {},
	}
	_, ok := m[i]
	return !ok
}

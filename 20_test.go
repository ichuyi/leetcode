package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("()[]"))
	fmt.Println(isValid("(]"))
}

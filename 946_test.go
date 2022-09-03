package main

import (
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

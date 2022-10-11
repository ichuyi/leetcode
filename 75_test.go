package main

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	fmt.Println(input)
}

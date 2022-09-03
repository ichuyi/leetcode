package main

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	fmt.Println(reversePrint(getLinkedList([]int{1, 2, 3, 4, 5})))
}

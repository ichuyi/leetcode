package main

import (
	"fmt"
	"testing"
)

func TestGetKthFromEnd(t *testing.T) {
	fmt.Println(getKthFromEnd(getLinkedList([]int{1, 2, 3, 4, 5}), 3))
}

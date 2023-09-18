package main

import (
	"testing"
)

func TestReorderList(t *testing.T) {
	root := getLinkedList([]int{1, 2, 3})
	reorderList(root)
	printLinkedList(root)
}

package main

import "testing"

func TestPartition(t *testing.T) {
	root := getLinkedList([]int{1, 2, 3})
	root = partition(root, 4)
	printLinkedList(root)
}

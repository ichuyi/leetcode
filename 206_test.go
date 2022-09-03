package main

import "testing"

func TestReverseList(t *testing.T) {
	root := getLinkedList([]int{1, 2, 3, 4, 5})
	root = reverseList2(root)
	printLinkedList(root)

}
func TestReverseList2(t *testing.T) {
	root := getLinkedList([]int{1, 2, 3, 4, 5})
	root = reverseList2(root)
	printLinkedList(root)

}

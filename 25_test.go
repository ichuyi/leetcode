package main

import "testing"

func TestReverseKGroup(t *testing.T) {
	printLinkedList(reverseKGroup(getLinkedList([]int{1, 2, 3, 4, 5}), 2))
}

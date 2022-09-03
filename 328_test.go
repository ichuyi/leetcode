package main

import "testing"

func TestOddEvenList(t *testing.T) {
	printLinkedList(oddEvenList(getLinkedList([]int{1, 2, 3, 4, 5})))
}

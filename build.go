package main

import "fmt"

func getLinkedList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	root := &ListNode{
		Val: input[0],
	}
	root.Next = getLinkedList(input[1:])
	return root
}
func printLinkedList(root *ListNode) {
	input := make([]int, 0)
	for s := root; s != nil; s = s.Next {
		input = append(input, s.Val)
	}
	fmt.Println(input)
}

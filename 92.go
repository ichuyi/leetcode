package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	root := &ListNode{
		Next: head,
	}
	firstEnd := root
	for i := 1; i < left; i++ {
		firstEnd = firstEnd.Next
	}
	secondEnd := firstEnd.Next
	for i := left; i < right; i++ {
		secondEnd = secondEnd.Next
	}
	thirdStart := secondEnd.Next
	secondEnd.Next = nil
	ss, se := reverse(firstEnd.Next)
	firstEnd.Next = ss
	se.Next = thirdStart
	return root.Next
}
func reverse(root *ListNode) (*ListNode, *ListNode) {
	start := root
	var current *ListNode
	last := root
	for current = root.Next; current != nil; {
		n := current.Next
		current.Next = last
		last = current
		current = n
	}
	return last, start
}

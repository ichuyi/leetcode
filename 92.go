package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	root := &ListNode{Next: head}
	start := root
	last := root
	for i := 1; i <= left; i++ {
		last = start
		start = start.Next
	}
	end := start
	for j := left + 1; j <= right; j++ {
		end = end.Next
	}
	next := end.Next
	end.Next = nil
	midStart, midEnd := reverse(start)
	last.Next = midStart
	midEnd.Next = next
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

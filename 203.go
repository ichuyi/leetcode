package main

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{}
	root.Next = head
	last := root
	for c := root.Next; c != nil; {
		if c.Val == val {
			n := c.Next
			last.Next = n
			c = n
		} else {
			last = c
			c = c.Next
		}
	}
	return root.Next
}

package main

func swapPairs(head *ListNode) *ListNode {
	root := &ListNode{Next: head}
	last := root
	for cur := head; cur != nil && cur.Next != nil; {
		next := cur.Next.Next
		last.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		last = cur
		cur = next
	}
	return root.Next
}

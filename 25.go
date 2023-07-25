package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{Next: head}
	index := 0
	lastEnd := root
	lastStart := root.Next
	for current := root.Next; current != nil; {
		index++
		if index%k == 0 {
			next := current.Next
			current.Next = nil
			s, e := reverse(lastStart)
			lastEnd.Next = s
			e.Next = next
			current = next
			lastStart = current
			lastEnd = e
		} else {
			current = current.Next
		}
	}
	return root.Next
}

func reverseItem(head *ListNode) (*ListNode, *ListNode) {
	var cur *ListNode
	var last *ListNode
	for cur = head; cur != nil; {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last, head
}

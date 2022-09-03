package main

func removeElements(head *ListNode, val int) *ListNode {
	var root *ListNode
	for root = head; root != nil && root.Val == val; root = root.Next {
	}
	last := root
	if root == nil {
		return nil
	}
	for start := root.Next; start != nil; {
		if start.Val == val {
			n := start.Next
			last.Next = n
			start = n
		} else {
			last = start
			start = start.Next
		}
	}
	return root
}

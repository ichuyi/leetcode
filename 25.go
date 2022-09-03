package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	var root *ListNode
	index := 0
	var groupStart, lastEnd *ListNode
	groupStart = head
	for current := head; current != nil; {
		index++
		if index%k == 0 {
			n := current.Next
			current.Next = nil
			s, e := reverseItem(groupStart)
			if root == nil {
				root = s
			} else {
				lastEnd.Next = s
			}
			lastEnd = e
			current = n
			groupStart = current
		} else {
			current = current.Next
		}
	}
	lastEnd.Next = groupStart
	return root
}
func reverseItem(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	var last *ListNode = head
	for start := head.Next; start != nil; {
		n := start.Next
		start.Next = last
		last = start
		start = n
	}
	head.Next = nil
	return last, head
}

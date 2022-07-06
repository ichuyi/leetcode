package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	var leftStart, leftEnd, midStart *ListNode
	if left > 1 {
		leftStart, leftEnd = head, head
		for i := 1; i < left-1; i++ {
			leftEnd = leftEnd.Next
		}
		midStart = leftEnd.Next
	} else {
		midStart = head
	}
	midEnd := midStart
	var last *ListNode
	var next *ListNode
	for i := left; i <= right; i++ {
		next = midStart.Next
		midStart.Next = last
		last = midStart
		midStart = next
	}
	if leftEnd != nil {
		leftEnd.Next = last
	} else {
		leftStart = last
	}
	midEnd.Next = next
	return leftStart
}

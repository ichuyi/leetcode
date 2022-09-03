package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var first, second *ListNode
	first = head
	second = head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}

package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	l := reverseList(head.Next)
	head.Next = nil
	n.Next = head
	return l
}
func reverseList2(head *ListNode) *ListNode {
	var last *ListNode
	var cur *ListNode
	for cur = head; cur != nil; {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}

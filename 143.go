package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	medium := findMedium(head)
	second := medium.Next
	medium.Next = nil
	second = reverseSecond(second)
	merge(head, second)
}
func findMedium(start *ListNode) *ListNode {
	fast, slow := start, start
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func reverseSecond(root *ListNode) *ListNode {
	var last, cur *ListNode
	for cur = root; cur != nil; {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}
func merge(first, second *ListNode) *ListNode {
	root := &ListNode{}
	last := root
	for first != nil && second != nil {
		firstNext, secondNext := first.Next, second.Next
		last.Next = first
		last.Next.Next = second
		last = second
		first = firstNext
		second = secondNext
	}
	last.Next = first
	return root.Next
}

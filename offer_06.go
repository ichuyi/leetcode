package main

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	r := make([]int, 0)
	reverseLinkedList(head, &r)
	return r
}
func reverseLinkedList(current *ListNode, r *[]int) {
	if current != nil {
		reverseLinkedList(current.Next, r)
	}
	*r = append(*r, current.Val)
}

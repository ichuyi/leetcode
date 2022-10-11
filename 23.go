package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var root *ListNode
	for i := 0; i < len(lists); i++ {
		root = mergeTwoList(root, lists[i])
	}
	return root
}
func mergeTwoList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var root = new(ListNode)
	last := root
	for right != nil && left != nil {
		if left.Val < right.Val {
			last.Next = left
			left = left.Next
			last = last.Next
		} else {
			last.Next = right
			right = right.Next
			last = last.Next
		}
	}
	if right != nil {
		last.Next = right
	} else {
		last.Next = left
	}
	return root.Next
}

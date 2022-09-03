package main

func oddEvenList(head *ListNode) *ListNode {
	var oddStart, oddEnd, evenStart, evenEnd *ListNode
	index := 0
	for current := head; current != nil; current = current.Next {
		index++
		if index%2 == 0 {
			if evenStart == nil {
				evenStart = current
			} else {
				evenEnd.Next = current
			}
			evenEnd = current
		} else {
			if oddStart == nil {
				oddStart = current
			} else {
				oddEnd.Next = current
			}
			oddEnd = current
		}
	}
	if oddEnd != nil {
		oddEnd.Next = evenStart
	}
	if evenEnd != nil {
		evenEnd.Next = nil
	}
	return oddStart
}

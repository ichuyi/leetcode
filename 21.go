package main

// getIntersectionNode TODO
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	return do(headA, headA, headB, headB)
}
func do(rootA, curA, rootB, curB *ListNode) *ListNode {
	if curA == nil && curB == nil {
		return nil
	}
	if curA == nil {
		curA = rootB
	}
	if curB == nil {
		curB = rootA
	}
	if curA == curB {
		return curB.Next
	}
	return do(rootA, curA.Next, rootB, curB.Next)
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	la, lb := 0, 0
	for s := headA; s != nil; s = s.Next {
		la++
	}
	for s := headB; s != nil; s = s.Next {
		lb++
	}
	sa, sb := headA, headB
	for i := 0; i < la-lb; i++ {
		sa = sa.Next
	}
	for i := 0; i < lb-la; i++ {
		sb = sb.Next
	}
	for {
		if sa == nil || sb == nil {
			return nil
		}
		if sa == sb {
			return sa
		}
		sa = sa.Next
		sb = sb.Next
	}
}

package main

// ListNode TODO
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
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

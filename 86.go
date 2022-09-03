package main

func partition(head *ListNode, x int) *ListNode {
	// 根节点  比x小的上一个节点 遍历的上一个节点
	root, lastMin, last := head, head, head
	if root == nil || root.Next == nil {
		return root
	}
	for cur := head.Next; cur != nil; {
		if cur.Val < x {
			// 如果根节点比x大，说明他不应该当根节点
			if root.Val >= x {
				root = cur
			}
			// 如果lastMin比x小，则把cur放到lastMin后面，并且更新lastMin
			n := cur.Next

			if lastMin.Val < x {
				// 这时候不用交换顺序
				if lastMin != last {
					last.Next = n
					next := lastMin.Next
					lastMin.Next = cur
					cur.Next = next
				} else {
					last = cur
				}
			} else {
				// 否则，cur应该放在lastMin前面，这时候其实他应该是根节点，但是根节点在前面更新过了，所以不用更新
				last.Next = n
				cur.Next = lastMin
			}

			lastMin = cur
			cur = n
		} else {
			last = cur
			cur = cur.Next
		}
	}
	return root
}

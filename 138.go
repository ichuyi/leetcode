package main

// Node TODO
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldNodeIndex := make(map[*Node]int)
	newIndexNode := make(map[int]*Node)
	var root, last *Node
	var i int
	for start := head; start != nil; start = start.Next {
		n := &Node{
			Val: start.Val,
		}
		if root == nil {
			root = n
		} else {
			last.Next = n
		}
		i++
		last = n
		oldNodeIndex[start] = i
		newIndexNode[i] = n
	}
	for start := head; start != nil; start = start.Next {
		rindex := oldNodeIndex[start]
		tindex := oldNodeIndex[start.Random]
		nr := newIndexNode[rindex]
		nt := newIndexNode[tindex]
		nr.Random = nt
	}
	return root
}
func copyRandomList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for start := head; start != nil; {
		next := start.Next
		n := &ListNode{
			Val: start.Val,
		}
		start.Next = n
		n.Next = next
		start = next
	}
	// for start := head; start != nil; start = start.Next.Next {
	// 	if start.Random != nil {
	// 		next := start.Next
	// 		next.Random = start.Random.Next
	// 	}
	// }
	root := head.Next
	for start := head; start != nil; start = start.Next {
		first := start.Next          // 2
		start.Next = start.Next.Next // 1->3
		if start.Next != nil {
			first.Next = start.Next.Next // 2->4
		}
	}
	printLinkedList(root)
	printLinkedList(head)
	return root
}
func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	for root := head; root != nil; root = root.Next {
		curCopy := &Node{
			Val: root.Val,
		}
		next := root.Next
		root.Next = curCopy
		curCopy.Next = next
		root = curCopy
	}
	for root := head; root != nil; root = root.Next {
		curCopy := root.Next
		if root.Random != nil {
			curCopy.Random = root.Random.Next
		}
		root = curCopy
	}
	root := head.Next
	curCopy := &Node{}
	for cur := head; cur != nil; cur = cur.Next {
		copyNext := cur.Next
		next := copyNext.Next
		cur.Next = next
		curCopy.Next = copyNext
		curCopy = copyNext
	}
	return root
}

package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[int]*TreeNode)
	recordParent(root, parent)
	m := make(map[int]struct{})
	ok := true
	cur := p
	for ok {
		m[cur.Val] = struct{}{}
		cur, ok = parent[cur.Val]
	}
	cur = q
	ok = true
	for ok {
		if _, o := m[cur.Val]; o {
			return cur
		}
		cur, ok = parent[cur.Val]
	}
	return nil
}
func recordParent(root *TreeNode, parent map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		parent[root.Left.Val] = root
	}
	if root.Right != nil {
		parent[root.Right.Val] = root
	}
	recordParent(root.Left, parent)
	recordParent(root.Right, parent)
}

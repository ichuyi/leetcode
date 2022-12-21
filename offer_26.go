package main
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if b == nil || a == nil {
			return false
	}
	if contains(a, b) {
			return true
	}
	return isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}

func contains(a, b *TreeNode) bool {
	if b == nil {
			return true
	}
	if a == nil {
			return false
	}
	if a.Val != b.Val {
			return false
	}
	return contains(a.Left, b.Left) && contains(a.Right, b.Right)
}
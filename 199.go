package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		var r []int
		r, cur = doLayer(cur)
		result = append(result, r[len(r)-1])
	}
	return result
}
func doLayer(input []*TreeNode) ([]int, []*TreeNode) {
	result := make([]int, 0)
	next := make([]*TreeNode, 0)
	for _, v := range input {
		result = append(result, v.Val)
		if v.Left != nil {
			next = append(next, v.Left)
		}
		if v.Right != nil {
			next = append(next, v.Right)
		}
	}
	return result, next
}

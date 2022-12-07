package main

// TreeNode TODO
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	d := true
	nextLayer := []*TreeNode{root}
	for len(nextLayer) > 0 {
		var t []int
		t, nextLayer = printLayer(nextLayer, d)
		result = append(result, t)
		d = !d
	}
	return result
}
func printLayer(input []*TreeNode, d bool) ([]int, []*TreeNode) {
	result := make([]int, 0)
	nextLayer := make([]*TreeNode, 0)
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i].Val)
		if d {
			if input[i].Left != nil {
				nextLayer = append(nextLayer, input[i].Left)
			}
			if input[i].Right != nil {
				nextLayer = append(nextLayer, input[i].Right)
			}
		} else {
			if input[i].Right != nil {
				nextLayer = append(nextLayer, input[i].Right)
			}
			if input[i].Left != nil {
				nextLayer = append(nextLayer, input[i].Left)
			}
		}
	}
	return result, nextLayer
}

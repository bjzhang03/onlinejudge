package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		result := &TreeNode{root.Val, nil, nil}
		result.Left = invertTree(root.Right)
		result.Right = invertTree(root.Left)
		return result
	}
	return nil
}

func main() {

}

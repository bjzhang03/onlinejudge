package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		if root.Left != nil {
			result = append(result, inorderTraversal(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, inorderTraversal(root.Right)...)
		}
	}
	return result
}

func main() {

}

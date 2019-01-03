package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	result := 0
	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth > rightDepth {
			result = leftDepth + 1
		} else {
			result = rightDepth + 1
		}
	}
	return result
}

func isBalanced(root *TreeNode) bool {
	result := true
	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if math.Abs(float64(leftDepth-rightDepth)) > 1 {
			result = false
		} else {
			result = result && isBalanced(root.Left) && isBalanced(root.Right)
		}
	}
	return result
}

func main() {

}

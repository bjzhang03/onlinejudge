package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 算法很差劲,但是过了
func countNodes(root *TreeNode) int {
	result := 0
	if root != nil {
		result = result + 1
		result = result + countNodes(root.Left) + countNodes(root.Right)
	}
	return result
}

func main() {

}

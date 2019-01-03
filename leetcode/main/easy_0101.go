package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symmetric(p *TreeNode, q *TreeNode) bool {
	result := true
	if p == nil && q == nil {
		result = true
	} else if p != nil && q == nil {
		result = false
	} else if p == nil && q != nil {
		result = false
	} else if p != nil && q != nil && p.Val != q.Val {
		result = false
	} else {
		result = result && symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
	}
	return result
}

func isSymmetric(root *TreeNode) bool {
	result := true
	if root != nil {
		result = symmetric(root.Left, root.Right)
	}
	return result
}

func main() {

}

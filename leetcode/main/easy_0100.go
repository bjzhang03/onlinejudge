package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	result := true
	if p != nil && q == nil {
		result = false
	} else if p == nil && q != nil {
		result = false
	} else if p != nil && q != nil && p.Val != q.Val {
		result = false
	} else if p == nil && q == nil {
		result = true
	} else if p != nil && q != nil && p.Val == q.Val {
		result = result && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return result
}

func main() {

}

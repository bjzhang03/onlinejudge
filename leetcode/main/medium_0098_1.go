package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://blog.csdn.net/zhangxiao93/article/details/50858927
func helper(root *TreeNode, left int, right int) bool {
	if root == nil {
		return true
	} else if left >= root.Val || right <= root.Val {
		return false
	}
	return helper(root.Left, left, root.Val) && helper(root.Right, root.Val, right)
}

func isValidBST(root *TreeNode) bool {
	left := math.MinInt64
	right := math.MaxInt64
	return helper(root, left, right)
}

func main() {
	fmt.Println(isValidBST(nil))

}

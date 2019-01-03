package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root != nil {
		result = [][]int{}
		queue := []*TreeNode{}
		currentLeval := []*TreeNode{}

		queue = append(queue, root)
		for len(queue) > 0 {
			currentLeval = []*TreeNode{}
			currentLevalValue := []int{}

			for len(queue) > 0 {
				tmp := queue[0]
				queue = queue[1:]
				currentLevalValue = append(currentLevalValue, tmp.Val)
				// 添加左子树
				if tmp.Left != nil {
					currentLeval = append(currentLeval, tmp.Left)
				}
				// 添加右子树
				if tmp.Right != nil {
					currentLeval = append(currentLeval, tmp.Right)
				}
			}
			result = append(result, currentLevalValue)
			queue = currentLeval
		}
	}

	res := [][]int{}
	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}
	return res
}
func main() {
	root := &TreeNode{3, nil, nil}
	fmt.Println(levelOrderBottom(root))
}

package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeSum struct {
	*TreeNode
	Sum int
}

func hasPathSum(root *TreeNode, sum int) bool {
	sumResult := []int{}
	result := false
	fmt.Println(sumResult)
	if root != nil {
		queue := []*TreeNodeSum{}
		queue = append(queue, &TreeNodeSum{root, root.Val})

		for len(queue) > 0 {
			tmp := queue[0]
			queue = queue[1:]

			if tmp.Left == nil && tmp.Right == nil {
				sumResult = append(sumResult, tmp.Sum)
			} else {
				if tmp.Left != nil {
					queue = append(queue, &TreeNodeSum{tmp.Left, tmp.Sum + tmp.Left.Val})
				}

				if tmp.Right != nil {
					queue = append(queue, &TreeNodeSum{tmp.Right, tmp.Sum + tmp.Right.Val})
				}
			}
		}
	}
	for i := 0; i < len(sumResult); i++ {
		if sumResult[i] == sum {
			result = true
		}
	}
	return result

}

func main() {
	root := &TreeNode{1, nil, nil}
	fmt.Println(hasPathSum(root, 1))

}

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeDepth struct {
	*TreeNode
	Depth int
}

func minDepth(root *TreeNode) int {
	result := 0
	if root != nil {
		queue := []*TreeNodeDepth{}
		rootDepth := &TreeNodeDepth{root, 1}
		queue = append(queue, rootDepth)
		for len(queue) > 0 {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Left == nil && tmp.Right == nil {
				result = tmp.Depth
				break
			} else {
				if tmp.Left != nil {
					queue = append(queue, &TreeNodeDepth{tmp.Left, tmp.Depth + 1})
				}
				if tmp.Right != nil {
					queue = append(queue, &TreeNodeDepth{tmp.Right, tmp.Depth + 1})
				}
			}
		}
	}
	return result
}

func main() {

}

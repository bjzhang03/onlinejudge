package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root != nil {
		parentsave := []*TreeNode{}
		parentsave = append(parentsave, root)
		for len(parentsave) > 0 {
			childsave := []*TreeNode{}
			value := []int{}
			for len(parentsave) > 0 {
				value = append(value, parentsave[0].Val)
				if parentsave[0].Left != nil {
					childsave = append(childsave, parentsave[0].Left)
				}
				if parentsave[0].Right != nil {
					childsave = append(childsave, parentsave[0].Right)
				}
				parentsave = parentsave[1:]
			}
			parentsave = []*TreeNode{}
			parentsave = append(parentsave, childsave...)
			result = append(result, value)
		}
	}
	return result
}

func main() {

}

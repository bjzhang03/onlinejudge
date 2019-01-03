package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		save := []*TreeNode{}
		save = append(save, root)

		for len(save) > 0 {
			result = append(result, save[0].Val)
			if save[0].Left != nil {
				save = append(save, save[0].Left)
			}

			if save[0].Right != nil {
				save = append(save, save[0].Right)
			}
			save = save[1:]
		}
	}

	//fmt.Println("result = ", result)
	return result
}

func isValidBST(root *TreeNode) bool {
	result := true
	if root != nil {
		// 如果左子树不是，则返回
		if !isValidBST(root.Left) {
			result = false
			return result
		}
		// 如果右子树不是则返回
		if !isValidBST(root.Right) {
			result = false
			return result
		}

		left := bfs(root.Left)
		right := bfs(root.Right)

		for i := 0; i < len(left); i++ {
			if left[i] >= root.Val {
				result = false
				return result
			}
		}

		for i := 0; i < len(right); i++ {
			if right[i] <= root.Val {
				result = false
				return result
			}
		}
	}
	return result
}

func main() {

}

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root != nil {
		tmpresult := [][]int{}
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
			tmpresult = append(tmpresult, value)
		}

		flag := true
		for i := 0; i < len(tmpresult); i++ {
			if !flag {
				tmp := []int{}
				for j := len(tmpresult[i]) - 1; j >= 0; j-- {
					tmp = append(tmp, tmpresult[i][j])
				}
				result = append(result, tmp)

			} else {
				result = append(result, tmpresult[i])
			}
			flag = !flag
		}
	}
	return result

}

func main() {

}

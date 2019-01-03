package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum int, result *[][]int, tmp []int) {
	if root.Left == nil && root.Right == nil {
		tmp = append(tmp, root.Val)
		tmpsum := 0
		for i := 0; i < len(tmp); i++ {
			tmpsum = tmpsum + tmp[i]
		}
		if tmpsum == sum {
			save := []int{}
			save = append(save, tmp...)
			*result = append(*result, save)
		}
		tmp = tmp[:len(tmp)-1]
	} else {
		if root.Left != nil {
			tmp = append(tmp, root.Val)
			dfs(root.Left, sum, result, tmp)
			tmp = tmp[:len(tmp)-1]
		}
		if root.Right != nil {
			tmp = append(tmp, root.Val)
			dfs(root.Right, sum, result, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	if root != nil {
		dfs(root, sum, &result, []int{})
	}
	return result
}

func main() {

}

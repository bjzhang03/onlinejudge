package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		level := []*TreeNode{}
		level = append(level, root)
		for len(level) > 0 {
			lcount := []int{}
			nlevel := []*TreeNode{}
			for len(level) > 0 {
				lcount = append(lcount, level[0].Val)
				if level[0].Left != nil {
					nlevel = append(nlevel, level[0].Left)
				}
				if level[0].Right != nil {
					nlevel = append(nlevel, level[0].Right)
				}
				level = level[1:]
			}
			result = append(result, lcount[len(lcount)-1])
			level = nlevel
		}
	}
	return result
}

func main() {

}

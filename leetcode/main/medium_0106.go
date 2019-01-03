package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) > 0 && len(postorder) > 0 {
		root := &TreeNode{postorder[len(postorder)-1], nil, nil}
		inoindex := len(postorder) - 1

		for inoindex > 0 {
			if inorder[inoindex] == postorder[len(postorder)-1] {
				break
			}
			inoindex--
		}
		//fmt.Println("inoindex = ", inoindex)
		leftpos := postorder[:inoindex]
		rightpos := postorder[inoindex : len(postorder)-1]
		leftino := inorder[:inoindex]
		rightino := inorder[inoindex+1:]
		//fmt.Println("leftpos = ", leftpos)
		//fmt.Println("rightpos = ", rightpos)
		//fmt.Println("leftino = ", leftino)
		//fmt.Println("rightino = ", rightino)
		root.Left = buildTree(leftino, leftpos)
		root.Right = buildTree(rightino, rightpos)
		return root

	}
	return nil
}

func main() {

	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(inorder, postorder))

}

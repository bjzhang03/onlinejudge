package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) > 0 && len(inorder) > 0 {
		root := &TreeNode{preorder[0], nil, nil}
		inoindex := 0
		for inoindex < len(inorder) {
			if inorder[inoindex] == preorder[0] {
				break
			}
			inoindex++
		}
		leftpre := preorder[1 : inoindex+1]
		rightpre := preorder[inoindex+1:]
		leftino := inorder[:inoindex]
		rightino := inorder[inoindex+1:]
		//fmt.Println("leftpre = ", leftpre)
		//fmt.Println("rightpre = ", rightpre)
		//fmt.Println("leftino = ", leftino)
		//fmt.Println("rightino = ", rightino)
		root.Left = buildTree(leftpre, leftino)
		root.Right = buildTree(rightpre, rightino)
		return root
	}
	return nil
}

func main() {

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))

}

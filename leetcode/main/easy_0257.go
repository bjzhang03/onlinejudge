package main

import (
	"fmt"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepfirstsearch(root *TreeNode, path string, result *[]string) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			tmp := (path + strconv.Itoa(root.Val))
			*result = append(*result, tmp)
			fmt.Println(path)
		}
		if root.Left != nil {
			tmp := path + strconv.Itoa(root.Val) + "->"
			deepfirstsearch(root.Left, tmp, result)
		}

		if root.Right != nil {
			tmp := path + strconv.Itoa(root.Val) + "->"
			deepfirstsearch(root.Right, tmp, result)
		}

	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root != nil {
		result := []string{}
		path := ""
		deepfirstsearch(root, path, &result)
		fmt.Println(result)
		return result
	}
	return nil
}

func main() {

}

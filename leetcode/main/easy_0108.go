package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) > 0 {
		tmpValue := nums[len(nums)/2]
		result:=&TreeNode{tmpValue,nil,nil}
		left:=nums[0:len(nums)/2]
		right:=nums[len(nums)/2+1:]
		result.Left=sortedArrayToBST(left)
		result.Right=sortedArrayToBST(right)
		return result
	}
	return nil
}

func main() {
	nums := []int{1}
	fmt.Println(sortedArrayToBST(nums))
}

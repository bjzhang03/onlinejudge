package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		// 计算一共有多少个节点的数据
		count := 0
		item := head
		for item != nil {
			count++
			item = item.Next
		}

		rotatek := k % count
		fmt.Println(rotatek)

		// 找到开始进行逆转的节点
		start := 0
		item = head
		for start < count-rotatek {
			item = item.Next
			start++
		}

		// 在result的后面添加数据
		it := result
		for item != nil {
			it.Next = &ListNode{item.Val, nil}
			it = it.Next
			item = item.Next
		}

		// 把剩下的数据添加进来
		start = 0
		item = head
		for start < count-rotatek {
			it.Next = &ListNode{item.Val, nil}
			it = it.Next
			item = item.Next
			start++
		}
	}
	return result.Next
}

func main() {

}

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil {
		// 计算一下一共有多少个节点
		item := head
		count := 0
		for item != nil {
			fmt.Println(item.Val, count)
			count++
			item = item.Next
		}
		// fmt.Println("count = ", count, " head = ", head.Val)
		// 添加前面的数据
		count = count - n
		result := &ListNode{0, nil}
		item2 := result
		item = head
		// fmt.Println("count --- ", count)
		for count > 0 {
			item2.Next = &ListNode{item.Val, nil}
			item2 = item2.Next
			item = item.Next
			count--
		}
		// 跳过需要删除的数据
		// fmt.Println("head value --- ", item.Val)
		item = item.Next
		for item != nil {
			item2.Next = &ListNode{item.Val, nil}
			item2 = item2.Next
			item = item.Next
		}
		return result.Next
	}
	return nil
}

func main() {
	head := &ListNode{0, nil}
	testdata := []int{1, 2, 3, 4, 5}
	item := head
	for i := 0; i < len(testdata); i++ {
		item.Next = &ListNode{testdata[i], nil}
		item = item.Next
	}
	//item = head.Next
	//for item != nil {
	//	fmt.Println(item.Val)
	//	item = item.Next
	//}
	removeNthFromEnd(head.Next, 2)
}

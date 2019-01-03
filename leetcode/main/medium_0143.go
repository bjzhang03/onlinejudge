package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 通过不了，指针没有发生改变
func reorderList(head *ListNode) {
	if head != nil {
		result := &ListNode{0, nil}
		// 计算节点个数
		count := 0
		item := head
		// 获取反序数据
		save := &ListNode{0, nil}
		for item != nil {
			tmp := &ListNode{item.Val, save.Next}
			save.Next = tmp
			item = item.Next
			count++
		}

		item1 := head
		item2 := save.Next
		tmpres := result
		for count > 0 && item1 != nil && item2 != nil {
			tmp1 := &ListNode{item1.Val, nil}
			tmpres.Next = tmp1
			count--
			item1 = item1.Next
			tmpres = tmpres.Next

			if count <= 0 {
				break
			}

			tmp2 := &ListNode{item2.Val, nil}
			tmpres.Next = tmp2
			count--
			item2 = item2.Next
			tmpres = tmpres.Next
		}
		// 输出结果来看一下，正确的
		itres := result.Next
		for itres != nil {
			fmt.Println(itres.Val)
			itres = itres.Next
		}
		// 这里需要这么进行特殊处理
		head.Next = result.Next.Next
	}

}

func main() {

}

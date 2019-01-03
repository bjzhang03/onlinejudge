package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		tmp := result
		item := head

		for item != nil {
			// 出现重复的数据
			if item.Next != nil && item.Val == item.Next.Val {
				for item.Next != nil && item.Val == item.Next.Val {
					item = item.Next
				}
			} else {
				tmp.Next = &ListNode{item.Val, nil}
				tmp = tmp.Next
			}
			item = item.Next
		}
	}
	return result.Next
}

func main() {

}

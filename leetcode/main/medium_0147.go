package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		item := head.Next
		result.Next = &ListNode{head.Val, nil}

		for item != nil {
			tmp := &ListNode{item.Val, nil}
			tmpres := result
			for tmpres.Next != nil {
				if tmpres.Next.Val >= item.Val {
					tmp.Next = tmpres.Next
					tmpres.Next = tmp
					break
				}
				tmpres = tmpres.Next
			}
			// 出现的数字需要添加到最后
			if tmpres.Next == nil {
				tmpres.Next = tmp
			}
			item = item.Next
		}
	}
	return result.Next
}

func main() {

}

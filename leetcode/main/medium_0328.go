package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		item := head
		it := result
		// 找到奇数节点的数据
		for item != nil {
			tmp := &ListNode{item.Val, nil}
			it.Next = tmp
			it = it.Next
			item = item.Next
			if item != nil {
				item = item.Next
			}
		}
		item = head.Next
		for item != nil {
			tmp := &ListNode{item.Val, nil}
			it.Next = tmp
			it = it.Next
			item = item.Next
			if item != nil {
				item = item.Next
			}
		}
	}
	return result.Next
}

func main() {

}

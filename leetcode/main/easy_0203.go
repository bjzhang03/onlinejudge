package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil {
		result := &ListNode{0, nil}
		it := result
		for head != nil {
			if head.Val == val {
				head = head.Next
			} else {
				it.Next = &ListNode{head.Val, nil}
				head = head.Next
				it = it.Next
			}
		}
		return result.Next
	}
	return nil
}

func main() {

}

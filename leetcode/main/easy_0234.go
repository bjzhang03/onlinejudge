package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	result := true
	if head != nil {
		it := head
		save := &ListNode{head.Val, nil}
		it = it.Next
		for it != nil {
			tmp := &ListNode{it.Val, nil}
			tmp.Next = save
			save = tmp
			it = it.Next
		}

		for head != nil && save != nil {
			if head.Val != save.Val {
				result = false
				break
			}
			head = head.Next
			save = save.Next
		}
	}
	return result
}

func main() {

}

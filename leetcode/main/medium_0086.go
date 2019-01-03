package main

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		right := &ListNode{0, nil}
		tmpres := result
		tmprig := right

		item := head

		for item != nil {
			if item.Val < x {
				tmpres.Next = &ListNode{item.Val, nil}
				tmpres = tmpres.Next
				item = item.Next
			} else {
				tmprig.Next = &ListNode{item.Val, nil}
				tmprig = tmprig.Next
				item = item.Next
			}
		}
		// 拼接两部分的数据
		tmpres.Next = right.Next
	}
	return result.Next
}

func main() {

}

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		count := 1
		item := head
		tmp := result
		// 前面部分的数据直接复制过来
		for count < m {
			tmp.Next = &ListNode{item.Val, nil}
			item = item.Next
			tmp = tmp.Next
			count++
		}
		// 中间部分的数据需要进行逆转
		reverse := &ListNode{0, nil}

		for count <= n {
			tmpsave := &ListNode{item.Val, reverse.Next}
			item = item.Next
			// 链表的头插法
			reverse.Next = tmpsave
			count++
		}

		//fmt.Println("count = ", count)
		tmprev := reverse.Next
		for tmprev.Next != nil {
			tmprev = tmprev.Next

		}
		for item != nil {
			//fmt.Println(item.Val)
			tmprev.Next = &ListNode{item.Val, nil}
			item = item.Next
			tmprev = tmprev.Next
		}
		// 拼接起来
		tmp.Next = reverse.Next
	}
	return result.Next
}
func main() {
	head := &ListNode{0, nil}
	item := head
	for i := 1; i <= 5; i++ {
		item.Next = &ListNode{i, nil}
		item = item.Next
	}

	reverseBetween(head.Next, 2, 4)

}

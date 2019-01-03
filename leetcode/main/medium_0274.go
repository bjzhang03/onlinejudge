package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	result := 0
	if len(citations) > 0 {
		// 先对数据进行排序
		sort.Ints(citations)
		//fmt.Println(citations)

		i := len(citations) - 1
		for ; i > 0; i-- {
			if citations[i] >= len(citations)-i && citations[i-1] <= len(citations)-i {
				result = len(citations) - i
				break
			}
		}
		if i == 0 && citations[0] >= len(citations) {
			result = len(citations)
		}
	}
	return result

}

func main() {

	nums := []int{1, 1}
	fmt.Println(hIndex(nums))

}

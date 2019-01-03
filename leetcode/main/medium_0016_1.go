package main

import (
	"fmt"
	"sort"
)
// https://blog.csdn.net/jaster_wisdom/article/details/80468931
func threeSumClosest(nums []int, target int) int {
	if len(nums) > 0 {
		// 对数据进行排序操作
		sort.Ints(nums)
	}
	return 0
}

func main() {
	nums := []int{0, 2, 1, -3}
	fmt.Println(threeSumClosest(nums, 1))
}

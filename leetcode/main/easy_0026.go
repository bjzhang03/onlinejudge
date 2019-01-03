package main

import "fmt"

//https://studygolang.com/articles/9743
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			result = i
			break
		}
	}
	if result < 0 {
		result = len(nums)
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))

}

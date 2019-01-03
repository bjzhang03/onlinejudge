package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	result := false
	if len(nums) > 0 && k > 0 {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j <= i+k && j < len(nums); j++ {
				//fmt.Println(nums[i],nums[j])
				if nums[i] == nums[j] {
					result = true
					break
				}
			}
		}
	}
	return result

}

func main() {
	nums := []int{99, 99}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) > 0 {
		start := 0
		end := len(nums) - 1

		for start <= end {
			mid := (start + end) / 2

			if nums[mid] == target {
				// 向前找
				rs := mid
				for rs-1 >= 0 && nums[rs] == nums[rs-1] {
					rs--
				}
				if rs > 1 {
					result[0] = rs
				} else {
					if nums[0] == nums[rs] {
						result[0] = 0
					} else {
						result[0] = 1
					}
				}

				// 向后找
				re := mid
				for re+1 <= len(nums)-1 && nums[re] == nums[re+1] {
					re++
				}
				if re < len(nums)-2 {
					result[1] = re
				} else {
					if nums[re] == nums[len(nums)-1] {
						result[1] = len(nums) - 1
					} else {
						result[1] = re
					}
				}
				break
			} else if nums[mid] > target {
				end = mid - 1
			} else if nums[mid] < target {
				start = mid + 1
			}
		}
	}
	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 6))

}

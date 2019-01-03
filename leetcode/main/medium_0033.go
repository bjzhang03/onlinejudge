package main

import "fmt"

// https://www.cnblogs.com/ganganloveu/p/4148554.html
func binarysearch(nums []int, target int, start int, end int) int {
	result := -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			result = mid
			break
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	return result

}

func search(nums []int, target int) int {
	result := -1
	if len(nums) > 0 {
		start := 0
		end := len(nums) - 1

		for start <= end {
			if nums[start] < nums[end] {
				result = binarysearch(nums, target, start, end)
				break
			} else {
				mid := (start + end) / 2

				if nums[mid] == target {
					result = mid
					break
				} else if nums[mid] >= nums[start] {
					// 前半部分是有序的
					// 判断数据是在前面还是在后面
					if nums[mid] <= target {
						start = mid + 1
					} else if nums[start] <= target {
						end = mid - 1
					} else if nums[start] > target {
						start = mid + 1
					}

				} else if nums[mid] < nums[start] {
					// 后半部分是有序的
					// 判断数据是在前面还是在后面
					if nums[mid] > target {
						end = mid - 1
					} else if target >= nums[start] {
						end = mid - 1
					} else if target < nums[start] {
						start = mid + 1
					}

				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	fmt.Println(search(nums, 9))

}

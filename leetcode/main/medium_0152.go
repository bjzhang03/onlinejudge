package main

import "fmt"

// https://www.cnblogs.com/bakari/p/4007368.html
func minInt(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {
	result := 0
	if len(nums) > 0 {
		if len(nums) == 1 {
			result = nums[0]
		} else {
			maxlocal := nums[0]
			minlocal := nums[0]
			result = nums[0]
			for i := 1; i < len(nums); i++ {
				maxcopy := maxlocal
				maxlocal = maxInt(maxInt(nums[i]*maxlocal, nums[i]), nums[i]*minlocal)
				minlocal = minInt(minInt(nums[i]*maxcopy, nums[i]), nums[i]*minlocal)
				result = maxInt(result, maxlocal)
			}
		}
	}
	return result
}

func main() {
	nums := []int{-4, -3, -2}
	fmt.Print(maxProduct(nums))
}

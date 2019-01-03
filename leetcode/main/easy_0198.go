package main

import "fmt"

func rob(nums []int) int {
	result := 0
	if len(nums) > 0 {
		if len(nums) == 1 {
			result = nums[0]
		} else if len(nums) == 2 {
			if nums[0] > nums[1] {
				result = nums[0]
			} else {
				result = nums[1]
			}
		} else if len(nums) == 3 {
			if nums[0]+nums[2] > nums[1] {
				result = nums[0] + nums[2]
			} else {
				result = nums[1]
			}
		} else {
			tmp1 := nums[0]
			tmp2 := 0
			if nums[0] > nums[1] {
				tmp2 = nums[0]
			} else {
				tmp2 = nums[1]
			}
			tmp3 := 0
			if nums[0]+nums[2] > nums[1] {
				tmp3 = nums[0] + nums[2]
			} else {
				tmp3 = nums[1]
			}

			for i := 3; i < len(nums); i++ {
				tmp := tmp3
				if tmp < tmp2+nums[i] {
					tmp = tmp2 + nums[i]
				}

				if tmp < tmp1+nums[i] {
					tmp = tmp1 + nums[i]
				}

				// 更新数据
				tmp1 = tmp2
				tmp2 = tmp3
				tmp3 = tmp
			}
			result = tmp3
		}

	}
	return result
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))

}

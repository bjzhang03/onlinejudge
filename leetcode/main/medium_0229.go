package main

import (
	"fmt"
	"math"
)

// https://blog.csdn.net/magicbean2/article/details/73737526
func majorityElement(nums []int) []int {
	result := []int{}
	if len(nums) > 0 {
		// 找到候选的两个数字
		acan := math.MaxInt64 // a candidate
		acount := 0
		bcan := math.MaxInt64 // b candidate
		bcount := 0
		for i := 0; i < len(nums); i++ {
			if acan == nums[i] {
				acount++
			} else if bcan == nums[i] {
				bcount++
			} else {
				if acount == 0 {
					acan = nums[i]
					acount++
				} else if bcount == 0 {
					bcan = nums[i]
					bcount++
				} else if acount > 0 && bcount > 0 {
					acount--
					bcount--
				}
			}
		}
		// 对候选的数字进行进一步的判断操作
		acount = 0
		bcount = 0
		for i := 0; i < len(nums); i++ {
			if acan == nums[i] {
				acount++
			}
			if bcan == nums[i] {
				bcount++
			}
		}
		if acount > len(nums)/3 {
			result = append(result, acan)
		}
		if bcount > len(nums)/3 {
			result = append(result, bcan)
		}
	}
	return result
}

func main() {

	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))

}

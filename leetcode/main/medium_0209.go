package main

import (
	"fmt"
	"math"
)

// https://blog.csdn.net/lisonglisonglisong/article/details/45666975
func minSubArrayLen(s int, nums []int) int {
	result := 0
	if len(nums) > 0 {
		ins := 0
		ine := 1
		minLength := math.MaxInt32
		sum := nums[0]
		for ins <= ine {
			//fmt.Println("ins = ", ins, ",ine = ", ine, ",sum = ", sum, ",minLength = ", minLength)
			tins := ins
			tine := ine
			if sum < s && ine < len(nums) {
				sum = sum + nums[ine]
				ine++
			} else if sum >= s && ins < len(nums) {
				if minLength > ine-ins {
					minLength = ine - ins
				}
				sum = sum - nums[ins]
				ins++
			}
			// 当数据的位置已经不再移动，则跳出来
			if tine == ine && tins == ins {
				break
			}
		}
		//fmt.Println(minLength, math.MaxInt32)
		if minLength != math.MaxInt32 {
			result = minLength
		}
	}
	return result
}

func main() {
	nums := []int{1, 1}
	fmt.Println("result = ", minSubArrayLen(7, nums))

}

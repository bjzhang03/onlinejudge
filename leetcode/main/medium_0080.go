package main

import "fmt"

func removeDuplicates(nums []int) int {
	result := 0
	if len(nums) > 0 {
		useless := []int{}
		start := 0
		for start < len(nums) {
			if start+1 < len(nums) {
				// 第一次出现重复的数据，不进行计算处理
				if nums[start] == nums[start+1] {
					start++
				}
				// 记下需要移除的数据的坐标
				for start+1 < len(nums) && nums[start] == nums[start+1] {
					useless = append(useless, start)
					start++
				}
			}
			start++
		}
		// fmt.Println(useless)
		for i := len(useless) - 1; i >= 0; i-- {
			for j := useless[i]; j < len(nums)-1; j++ {
				//fmt.Println("j = ", j)
				nums[j] = nums[j+1]
			}
			// fmt.Println("nums = ", nums)
		}
		result = len(nums) - len(useless)
		// fmt.Println(nums)
	}
	return result
}

func main() {

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))

}

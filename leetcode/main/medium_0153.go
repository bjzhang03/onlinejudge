package main

import "fmt"

func findMin(nums []int) int {
	result := 0
	if len(nums) == 1 {
		result = nums[0]
	} else if len(nums) > 1 {
		ins := 0
		ine := len(nums) - 1
		if nums[ine] > nums[ins] {
			// 此时是有序数组不需要查找
			result = nums[ins]
		} else {
			for ins < ine {
				//fmt.Println("ins = ", ins, ",ine = ", ine)
				if ins+1 == ine {
					result = nums[ine]
					break
				}

				inmid := (ins + ine) / 2

				if nums[inmid] < nums[ine] {
					ine = inmid
				} else if nums[inmid] > nums[ine] {
					ins = inmid
				}
			}
		}
	}
	return result
}

func main() {

	nums := []int{0, 1}
	fmt.Println(findMin(nums))

}

package main

import "fmt"

func lengthOfLIS(nums []int) int {
	result := 0
	if len(nums) > 0 {
		save := make([]int, len(nums))

		for i := 0; i < len(nums); i++ {
			tmp := 1
			for j := 0; j < i; j++ {
				if nums[j] < nums[i] && tmp < save[j]+1 {
					tmp = save[j] + 1
				}
			}
			save[i] = tmp

			if result < tmp {
				result = tmp
			}
		}
	}
	return result
}

func main() {

	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

}

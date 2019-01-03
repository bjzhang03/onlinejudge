package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) > 0 {
		save := make([]bool, len(nums))
		save[0] = true
		for i := 0; i < len(nums)-1; i++ {
			if save[i] {
				for j := 1; j <= nums[i]; j++ {
					if i+j < len(save) {
						save[i+j] = true
					}
				}
			}
			// fmt.Println("i = ", i)
			if save[len(save)-1] {
				return true
			}
		}
		if save[len(save)-1] {
			return true
		}
		// fmt.Println(save)
	}
	return false
}

func main() {
	nums := []int{0}
	fmt.Println(canJump(nums))

}

package main

import "fmt"

func moveZeroes(nums []int) {
	count := 0
	// 从后面的数据开始复制操作，避免出现0被多次复制的操作
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j+1 < len(nums); j++ {
				nums[j] = nums[j+1]
			}
			count++
		}
	}
	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

}

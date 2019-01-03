package main

import "fmt"

func majorityElement(nums []int) int {
	result := 0
	if len(nums) > 0 {
		stack := []int{}

		for i := 0; i < len(nums); i++ {
			if len(stack) == 0 {
				stack = append(stack, nums[i])
			} else if stack[len(stack)-1] == nums[i] {
				stack = append(stack, nums[i])
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		result = stack[0]
	}
	return result
}

func main() {
	nums:=[]int{2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))

}

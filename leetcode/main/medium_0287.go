package main

// 一个比较难以理解的方法
// 在leetcode上面看到的解法
// 看了很久还是不太明白啊
func findDuplicate(nums []int) int {
	result := 0
	if len(nums) > 0 {
		// Find the intersection point of the two runners.
		tortoise := nums[0]
		hare := nums[0]

		tortoise = nums[tortoise]
		hare = nums[nums[hare]]

		for tortoise != hare {
			tortoise = nums[tortoise]
			hare = nums[nums[hare]]
		}

		// Find the "entrance" to the cycle.
		ptr1 := nums[0]
		ptr2 := tortoise

		for ptr1 != ptr2 {
			ptr1 = nums[ptr1]
			ptr2 = nums[ptr2]
		}
		result = ptr1
	}
	return result
}

func main() {

}

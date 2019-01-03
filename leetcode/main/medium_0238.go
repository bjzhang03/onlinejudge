package main

func productExceptSelf(nums []int) []int {
	if len(nums) > 0 {
		result := make([]int, len(nums))
		mul := 1
		zeroCount := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] != 0 {
				mul = mul * nums[i]
			} else {
				zeroCount++
			}
		}
		if zeroCount == 0 {
			for i := len(nums) - 1; i >= 0; i-- {
				if nums[i] != 0 {
					result[i] = mul / nums[i]
				} else {
					result[i] = mul
				}
			}
		} else if zeroCount == 1 {
			for i := len(nums) - 1; i >= 0; i-- {
				if nums[i] == 0 {
					result[i] = mul
				}
			}
		}
		return result
	}
	return []int{}
}

func main() {

}

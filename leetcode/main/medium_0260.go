package main

// https://blog.csdn.net/u014673347/article/details/47728341
// https://blog.csdn.net/qq_38595487/article/details/81163737
func singleNumber(nums []int) []int {
	result := []int{}
	if len(nums) > 0 {
		save := make(map[int]bool)

		for i := 0; i < len(nums); i++ {
			if _, ok := save[nums[i]]; !ok {
				save[nums[i]] = true
			} else {
				delete(save, nums[i])
			}
		}

		for k, _ := range save {
			result = append(result, k)
		}
	}
	return result
}

func main() {

}

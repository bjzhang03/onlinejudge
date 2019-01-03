package main

func containsDuplicate(nums []int) bool {
	result := false
	if len(nums) > 0 {
		save := make(map[int]bool)

		for i := 0; i < len(nums); i++ {
			if _, ok := save[nums[i]]; ok {
				result = true
				break
			} else {
				save[nums[i]] = true
			}
		}
	}

	return result
}

func main() {

}

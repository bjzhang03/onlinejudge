package main

import "fmt"

func dfs(nums []int, result *[][]int, save []int) {
	// fmt.Println("nums = ", nums, "save = ", save)
	if len(nums) == 0 {
		// fmt.Println("if")
		tmp := []int{}
		tmp = append(tmp, save...)
		*result = append(*result, tmp)
		// fmt.Println("result = ", *result, "save = ", save)
		return
	} else if len(nums) > 0 {
		// fmt.Println("else")
		for i := 0; i < len(nums); i++ {
			// fmt.Println("nums = ", nums, " i = ", i)
			save = append(save, nums[i])
			nextnums := []int{}
			for j := 0; j < len(nums); j++ {
				if j != i {
					nextnums = append(nextnums, nums[j])
				}
			}
			dfs(nextnums, result, save)
			save = save[:len(save)-1]
			// fmt.Println("save = ", save)
		}

	}

}
func permute(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		save := []int{}
		dfs(nums, &result, save)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums[:0])
	fmt.Println(nums[1:])
	fmt.Println(permute(nums))

}

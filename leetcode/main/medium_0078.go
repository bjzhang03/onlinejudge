package main

import "fmt"

func dfs(nums []int, result *[][]int, tmp []int, index int) {
	// 将结果添加进来
	item := []int{}
	item = append(item, tmp...)
	*result = append(*result, item)
	if index < len(nums) {
		for i := index; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			dfs(nums, result, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		dfs(nums, &result, []int{}, 0)
	}
	return result
}

func main() {

	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))

}

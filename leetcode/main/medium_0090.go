package main

import (
	"fmt"
	"sort"
)

func hasslice(save [][]int, tmp []int) bool {
	result := false
	for i := 0; i < len(save); i++ {
		if len(save[i]) == len(tmp) {
			flag := true
			for j := 0; j < len(tmp); j++ {
				if tmp[j] != save[i][j] {
					flag = false
					break
				}
			}
			if flag {
				result = true
				break
			}
		}
	}
	return result
}

func dfs(result *[][]int, nums []int, index int, tmp []int) {
	if !hasslice(*result, tmp) {
		item := []int{}
		item = append(item, tmp...)
		*result = append(*result, item)
	}
	if index < len(nums) {
		for i := index; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			dfs(result, nums, i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		sort.Ints(nums)
		dfs(&result, nums, 0, []int{})
	}
	return result
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

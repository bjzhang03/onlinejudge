package main

import (
	"fmt"
	"sort"
)

func hasslice(save *[][]int, tmpsave []int) bool {
	result := false
	for i := 0; i < len(*save); i++ {
		if (*save)[i][0] == tmpsave[0] && (*save)[i][1] == tmpsave[1] && (*save)[i][2] == tmpsave[2] && (*save)[i][3] == tmpsave[3] {
			result = true
			break
		}
	}
	return result
}
func dfs(nums []int, result *[][]int, tmpsave []int, index int, currentsum int, target int) {
	// 已经存储了4个元素了的情况
	if len(tmpsave) == 4 {
		if target == tmpsave[0]+tmpsave[1]+tmpsave[2]+tmpsave[3] {
			if !hasslice(result, tmpsave) {
				*result = append(*result, []int{tmpsave[0], tmpsave[1], tmpsave[2], tmpsave[3]})
			}
		}
		return
	}
	for i := index; i < len(nums); i++ {
		nextcurrentsum := currentsum + nums[i]
		nexttmpsave := append(tmpsave, nums[i])
		dfs(nums, result, nexttmpsave, i+1, nextcurrentsum, target)
	}
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) >= 4 {
		// 先进行排序操作
		sort.Ints(nums)
		for i := 0; i < len(nums); i++ {
			currentsum := nums[i]
			tmpsave := []int{nums[i]}
			dfs(nums, &result, tmpsave, i+1, currentsum, target)
		}
	}
	return result
}

func main() {

	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	fmt.Println(fourSum(nums, -11))

}

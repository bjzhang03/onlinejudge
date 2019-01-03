package main

import (
	"fmt"
	"sort"
)

func hasslice(save [][]int, item []int) bool {
	result := false
	for i := 0; i < len(save); i++ {
		if len(save[i]) == len(item) {
			flag := true
			for j := 0; j < len(item); j++ {
				if save[i][j] != item[j] {
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

func dfs(candidates []int, target int, result *[][]int, save []int) {
	if target == 0 {
		// 最后的结果进行排序操作
		sort.Ints(save)
		// 检查结果是否已经出现过
		if !hasslice(*result, save) {
			// 这里需要新建一个数组，不然会出现错误
			tmp := []int{}
			tmp = append(tmp, save...)
			*result = append(*result, tmp)
			// fmt.Println("result = ", *result)
		}
		return
	} else if target > 0 {
		for i := 0; i < len(candidates); i++ {
			// 新建一个变量来处理
			nextsave := []int{}
			nextsave = append(nextsave, save...)
			nextsave = append(nextsave, candidates[i])
			dfs(candidates, target-candidates[i], result, nextsave)
		}
	}
}
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) > 0 {
		// 先对数据进行排序
		sort.Ints(candidates)
		save := []int{}
		dfs(candidates, target, &result, save)
	}
	return result
}

func main() {
	nums := []int{2, 3, 7}
	fmt.Println(combinationSum(nums, 18))

}

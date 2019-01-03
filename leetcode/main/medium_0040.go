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

func dfs(candidate []int, target int, result *[][]int, save []int, index int) {
	if target == 0 {
		sort.Ints(save)
		if !hasslice(*result, save) {
			tmp := []int{}
			tmp = append(tmp, save...)
			*result = append(*result, tmp)
			// fmt.Println("result = ", *result)
		}
		return
	} else if target > 0 {
		for i := index; i < len(candidate); i++ {
			save = append(save, candidate[i])
			dfs(candidate, target-candidate[i], result, save, i+1)
			save = save[:len(save)-1]
		}
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) > 0 {
		sort.Ints(candidates)
		save := []int{}
		dfs(candidates, target, &result, save, 0)
	}
	return result
}

func main() {
	candidate := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidate, 8))

}

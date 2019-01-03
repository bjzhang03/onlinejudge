package main

import "fmt"

func dfs(save []int, result *[][]int, tmp []int, index int, count int) {
	if len(tmp) == count {
		item := []int{}
		item = append(item, tmp...)
		*result = append(*result, item)
	} else if len(tmp) < count {
		for i := index; i < len(save); i++ {
			tmp = append(tmp, save[i])
			dfs(save, result, tmp, i+1, count)
			tmp = tmp[:len(tmp)-1]
		}
	}

}

func combine(n int, k int) [][]int {
	result := [][]int{}
	if n > 0 && k > 0 && n >= k {
		save := make([]int, n)
		for i := 0; i < n; i++ {
			save[i] = i + 1
		}
		// fmt.Println(save)
		dfs(save, &result, []int{}, 0, k)
	}
	return result
}

func main() {
	fmt.Println(combine(4, 2))

}

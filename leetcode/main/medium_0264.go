package main

import (
	"fmt"
	"sort"
)

func nthUglyNumber(n int) int {
	result := 0
	if n > 0 {
		candidates := []int{1}
		used := map[int]bool{1: true}
		save := []int{}
		for len(save) <= n {
			fmt.Println(save, candidates)
			tmpmul := candidates[0]
			candidates = candidates[1:]

			if _, ok := used[tmpmul*2]; !ok {
				candidates = append(candidates, tmpmul*2)
				used[tmpmul*2] = true
			}
			if _, ok := used[tmpmul*3]; !ok {
				candidates = append(candidates, tmpmul*3)
				used[tmpmul*3] = true
			}
			if _, ok := used[tmpmul*5]; !ok {
				candidates = append(candidates, tmpmul*5)
				used[tmpmul*5] = true
			}
			delete(used, tmpmul)
			// 添加一个数据进来
			save = append(save, tmpmul)
			sort.Ints(candidates)
		}
		result = save[n-1]
	}
	return result
}

func main() {

	fmt.Println(nthUglyNumber(100))

}

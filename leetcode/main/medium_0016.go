package main

import (
	"fmt"
	"math"
)
// 回溯的方法实现
func dfs(nums []int, countsave []int, index int, target int, result *int) {
	//fmt.Println("index = ", index, "result = ", *result, " countsave = ", countsave)
	// 临界判断
	if len(countsave) == 3 {
		if math.Abs(float64(target-*result)) > math.Abs(float64(target-(countsave[0]+countsave[1]+countsave[2]))) {
			*result = countsave[0] + countsave[1] + countsave[2]
		}
		return
	}
	// 添加元素进行处理
	for i := index; i < len(nums); i++ {
		tmp := append(countsave, nums[i])
		dfs(nums, tmp, i+1, target, result)
	}
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) > 0 {
		result := math.MaxInt32
		countsave := []int{}
		index := 0
		dfs(nums, countsave, index, target, &result)
		return result
	}
	return 0
}

func main() {

	nums := []int{0, 2, 1, -3}
	fmt.Println(threeSumClosest(nums, 1))

}

package main

import "fmt"

func sumInt(vals ...int) int {
	result := 0
	for i := 0; i < len(vals); i++ {
		result = result + vals[i]
	}
	return result
}

func combination(result *[][]int, start int, save []int, k int, n int) {
	if len(save) < k {
		for i := start; i <= 9; i++ {
			save = append(save, i)
			combination(result, i+1, save, k, n)
			save = save[:len(save)-1]
		}

	} else if len(save) == k {
		if sumInt(save...) == n {
			tmp := []int{}
			tmp = append(tmp, save...)
			*result = append(*result, tmp)
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	combination(&result, 1, []int{}, k, n)
	return result
}

func main() {

	fmt.Println(combinationSum3(3, 9))

}

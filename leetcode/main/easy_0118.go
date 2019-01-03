package main

import "fmt"

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows > 0 {
		tmp := []int{1}
		result = append(result, tmp)
		for i := 1; i < numRows; i++ {
			tmp = make([]int, i+1)
			tmp[0] = 1
			tmp[i] = 1
			for j := 1; j < i; j++ {
				tmp[j] = result[i-1][j-1] + result[i-1][j]
			}
			result = append(result, tmp)
		}
	}
	return result
}

func main() {
	fmt.Println(generate(5))

}

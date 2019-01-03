package main

import "fmt"

func getRow(rowIndex int) []int {
	result := []int{}
	if rowIndex >= 0 {
		result = []int{1}
		for i := 1; i <= rowIndex; i++ {
			tmp := make([]int, i+1)
			tmp[0] = 1
			tmp[i] = 1
			for j := 1; j < i; j++ {
				tmp[j] = result[j-1] + result[j]
			}
			result = tmp
		}
	}
	return result

}

func main() {

	fmt.Println(getRow(3))

}

package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	result := math.MaxInt32
	if len(triangle) > 0 {
		save := make([][]int, len(triangle))
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(triangle[i]))
		}
		save[0][0] = triangle[0][0]
		for i := 1; i < len(save); i++ {
			save[i][0] = save[i-1][0] + triangle[i][0]
			save[i][len(save[i])-1] = save[i-1][len(save[i-1])-1] + triangle[i][len(triangle[i])-1]
		}
		for i := 2; i < len(save); i++ {
			for j := 1; j < len(save[i])-1; j++ {
				if save[i-1][j-1] > save[i-1][j] {
					save[i][j] = save[i-1][j] + triangle[i][j]
				} else {
					save[i][j] = save[i-1][j-1] + triangle[i][j]
				}
			}
		}
		//fmt.Println(save)
		for i := 0; i < len(save[len(save)-1]); i++ {
			if result > save[len(save)-1][i] {
				result = save[len(save)-1][i]
			}
		}
	}
	return result
}

func main() {

	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 9, 8, 3}}
	fmt.Println(minimumTotal(triangle))

}

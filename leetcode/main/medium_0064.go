package main

import "fmt"

func minPathSum(grid [][]int) int {
	result := 0
	if len(grid) > 0 {
		rows := len(grid)
		cols := len(grid[0])

		save := make([][]int, rows)
		for i := 0; i < rows; i++ {
			save[i] = make([]int, cols)
		}

		// 行数据的初始化
		save[0][0] = grid[0][0]
		for i := 1; i < rows; i++ {
			save[i][0] = save[i-1][0] + grid[i][0]
		}
		// 列数据的初始化
		for i := 1; i < cols; i++ {
			save[0][i] = save[0][i-1] + grid[0][i]
		}

		for i := 1; i < rows; i++ {
			for j := 1; j < cols; j++ {
				if save[i-1][j] > save[i][j-1] {
					save[i][j] = save[i][j-1] + grid[i][j]
				} else {
					save[i][j] = save[i-1][j] + grid[i][j]
				}
			}
		}
		// fmt.Println(save)
		result = save[rows-1][cols-1]
	}
	return result
}

func main() {
	nums := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(nums))

}

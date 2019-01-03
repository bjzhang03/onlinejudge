package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	result := 0
	if len(obstacleGrid) > 0 {
		rows := len(obstacleGrid)
		cols := len(obstacleGrid[0])

		save := make([][]int, rows)
		for i := 0; i < rows; i++ {
			save[i] = make([]int, cols)
		}

		// 行方向的初始化
		for i := 0; i < cols; i++ {
			if obstacleGrid[0][i] == 1 {
				break
			} else {
				save[0][i] = 1
			}
		}
		// 列方向的初始化

		for i := 0; i < rows; i++ {
			if obstacleGrid[i][0] == 1 {
				break
			} else {
				save[i][0] = 1
			}
		}

		for i := 1; i < rows; i++ {
			for j := 1; j < cols; j++ {
				if obstacleGrid[i][j] != 1 {
					save[i][j] = save[i-1][j] + save[i][j-1]
				}
			}
		}
		// fmt.Println(save)
		result = save[rows-1][cols-1]
	}
	return result
}

func main() {
	nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(nums))

}

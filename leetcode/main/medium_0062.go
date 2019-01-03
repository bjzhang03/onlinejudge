package main

import "fmt"

func uniquePaths(m int, n int) int {
	result := 0
	if m > 0 && n > 0 {
		save := make([][]int, n)
		for i := 0; i < n; i++ {
			save[i] = make([]int, m)
		}

		for i := 0; i < n; i++ {
			save[i][0] = 1
		}
		for i := 0; i < m; i++ {
			save[0][i] = 1
		}

		for i := 1; i < n; i++ {
			for j := 1; j < m; j++ {
				save[i][j] = save[i-1][j] + save[i][j-1]
			}
		}
		// fmt.Println(save)
		result = save[n-1][m-1]
	}
	return result
}

func main() {

	fmt.Println(uniquePaths(10, 3))

}

package main

func setZeroes(matrix [][]int) {
	if len(matrix) > 0 {
		rows := make(map[int]bool)
		cols := make(map[int]bool)

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] == 0 {
					rows[i] = true
					cols[j] = true
				}
			}
		}

		for col, _ := range cols {
			for i := 0; i < len(matrix); i++ {
				matrix[i][col] = 0
			}
		}

		for row, _ := range rows {
			for i := 0; i < len(matrix[0]); i++ {
				matrix[row][i] = 0
			}
		}
		// fmt.Println(matrix)
	}

}

func main() {

	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(nums)

}

package main

func simplerotate(matrix [][]int, s int, e int) {
	// s start, e end
	if s < e {
		save := []int{}
		// 上
		for i := s; i < e; i++ {
			save = append(save, matrix[s][i])
		}
		// 右
		for i := s; i < e; i++ {
			save = append(save, matrix[i][e])
		}
		// 下
		for i := e; i > s; i-- {
			save = append(save, matrix[e][i])
		}
		// 左
		for i := e; i > s; i-- {
			save = append(save, matrix[i][s])
		}
		// fmt.Println(save)
		start := 0
		// 右
		for i := s; i < e; i++ {
			matrix[i][e] = save[start]
			start++
		}
		// 下
		for i := e; i > s; i-- {
			matrix[e][i] = save[start]
			start++
		}
		// 左
		for i := e; i > s; i-- {
			matrix[i][s] = save[start]
			start++
		}
		// 上
		for i := s; i < e; i++ {
			matrix[s][i] = save[start]
			start++
		}
		// fmt.Println(matrix)
	}

}
func rotate(matrix [][]int) {
	if len(matrix) > 0 {
		start := 0
		end := len(matrix) - 1
		for start <= end {
			simplerotate(matrix, start, end)
			start++
			end--
		}
	}

}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	simplerotate(matrix, 0, 2)

}

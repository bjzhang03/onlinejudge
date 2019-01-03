package main

import "fmt"

func simpleGenerateMatrix(matrix *[][]int, save []int, start *int, s int, e int) {
	if s < e {
		// 上
		for i := s; i < e; i++ {
			(*matrix)[s][i] = save[*start]
			(*start)++
		}
		// 右
		for i := s; i < e; i++ {
			(*matrix)[i][e] = save[*start]
			(*start)++
		}
		// 下
		for i := e; i > s; i-- {
			(*matrix)[e][i] = save[*start]
			(*start)++
		}
		// 左
		for i := e; i > s; i-- {
			(*matrix)[i][s] = save[*start]
			(*start)++
		}
	} else if s == e {
		(*matrix)[s][e] = save[*start]
		(*start)++
	}

}

func generateMatrix(n int) [][]int {
	if n > 0 {
		// 生成矩阵
		result := make([][]int, n)
		for i := 0; i < n; i++ {
			result[i] = make([]int, n)
		}
		// 生成数据
		save := make([]int, n*n)
		for i := 0; i < n*n; i++ {
			save[i] = i + 1
		}

		index := 0
		start := 0
		end := n - 1
		for start <= end {
			simpleGenerateMatrix(&result, save, &index, start, end)
			start++
			end--
		}
		return result

	}
	return [][]int{}
}

func main() {

	fmt.Println(generateMatrix(5))

}

package main

import "fmt"

func simpleSpirOrderer(matrix [][]int, cs int, ce int, rs int, re int) []int {
	result := []int{}
	if ce > cs && re > rs {
		// 上
		for i := cs; i < ce; i++ {
			result = append(result, matrix[rs][i])
		}
		// 右
		for i := rs; i < re; i++ {
			result = append(result, matrix[i][ce])
		}
		// 下
		for i := ce; i > cs; i-- {
			result = append(result, matrix[re][i])
		}
		// 左
		for i := re; i > rs; i-- {
			result = append(result, matrix[i][cs])
		}

	} else if ce > cs && re == rs {
		for i := cs; i <= ce; i++ {
			result = append(result, matrix[re][i])
		}
	} else if re > rs && ce == cs {
		for i := rs; i <= re; i++ {
			result = append(result, matrix[i][ce])
		}
	} else if re == rs && ce == cs {
		result = append(result, matrix[re][ce])
	}
	return result
}

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) > 0 {
		rows := 0
		cols := 0
		rowe := len(matrix) - 1
		cole := len(matrix[0]) - 1
		for rowe >= rows && cole >= cols {
			result = append(result, simpleSpirOrderer(matrix, cols, cole, rows, rowe)...)
			rows++
			rowe--
			cols++
			cole--
		}
	}
	return result
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))

}

package main

import (
	"fmt"
)

func isSquare(matrix [][]byte, row int, col int, length int) bool {
	result := false
	if row+length < len(matrix) && col+length < len(matrix[0]) {
		result = true
		for i := row; i <= row+length; i++ {
			for j := col; j <= col+length; j++ {
				//fmt.Println("i = ", i, "j = ", j)
				if matrix[i][j] != 1 {
					result = false
					break
				}
			}
		}
	}
	//fmt.Println(row, col, length, result)
	return result
}

func maximalSquare(matrix [][]byte) int {
	result := 0
	if len(matrix) > 0 {
		for length := 0; length < len(matrix); length++ {
			flag := false
			for row := 0; row < len(matrix); row++ {
				for col := 0; col < len(matrix[0]); col++ {
					if isSquare(matrix, row, col, length) {
						//fmt.Println(j, k, i)
						flag = true
						break
					}
				}
				// 如果找到了就不在继续找
				if flag {
					break
				}
			}
			if flag && result < (length+1)*(length+1) {
				result = (length + 1) * (length + 1)
			}
			//fmt.Println("length = ", length, " flag = ", flag, "result = ", result)
		}
	}
	return result
}

func main() {
	matrix := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	//matrix := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	fmt.Println(maximalSquare(matrix))
}

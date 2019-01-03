package main

import "fmt"

func binarySearchRows(matrix [][]int, target int, row int, cols int, cole int) bool {
	result := false
	for cols <= cole {
		colmid := (cols + cole) / 2
		if matrix[row][colmid] == target {
			result = true
			break
		} else if matrix[row][colmid] > target {
			cole = colmid - 1
		} else {
			cols = colmid + 1
		}
	}
	return result
}

func binarySearchCols(matrix [][]int, target int, col int, rows int, rowe int) bool {
	result := false
	for rows <= rowe {
		rowmid := (rows + rowe) / 2
		if matrix[rowmid][col] == target {
			result = true
			break
		} else if matrix[rowmid][col] > target {
			rowe = rowmid - 1
		} else {
			rows = rowmid + 1
		}
	}
	return result
}

func binarySearchMatrix(matrix [][]int, target int, rows int, rowe int, cols int, cole int) bool {
	// fmt.Println("rows = ", rows, " rowe = ", rowe, " cols = ", cols, " cole = ", cole)
	result := false
	if rowe >= rows && cole >= cols {
		rowmid := (rows + rowe) / 2
		colmid := (cols + cole) / 2
		if matrix[rowmid][colmid] == target {
			// fmt.Println("if")
			result = true
		} else if matrix[rowmid][colmid] > target {
			// fmt.Println("else if")
			if binarySearchCols(matrix, target, colmid, rows, rowmid) {
				result = true
			} else if binarySearchRows(matrix, target, rowmid, cols, colmid) {
				result = true
			} else {
				result = binarySearchMatrix(matrix, target, rows, rowmid-1, cols, colmid-1) || binarySearchMatrix(matrix, target, rows, rowmid-1, colmid+1, cole)
			}
		} else {
			// fmt.Println("else")
			if binarySearchRows(matrix, target, rowmid, colmid, cole) {
				result = true
			} else if binarySearchCols(matrix, target, colmid, rowmid, rowe) {
				result = true
			} else {
				result = binarySearchMatrix(matrix, target, rowmid+1, rowe, cols, colmid-1) || binarySearchMatrix(matrix, target, rowmid+1, rowe, colmid+1, cole)
			}
		}
	}
	return result

}

func searchMatrix(matrix [][]int, target int) bool {
	result := false
	if len(matrix) > 0 {
		rows := 0
		cols := 0
		rowe := len(matrix) - 1
		cole := len(matrix[0]) - 1
		result = binarySearchMatrix(matrix, target, rows, rowe, cols, cole)
	}
	return result
}

func main() {

	nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(searchMatrix(nums, 23))

}

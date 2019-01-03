package main

import "fmt"

type NumMatrix struct {
	save [][]int
	sum  [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	obj := &NumMatrix{[][]int{}, [][]int{}}
	obj.save = make([][]int, len(matrix))
	obj.sum = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		obj.save[i] = make([]int, len(matrix[0]))
		obj.sum[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			obj.save[i][j] = matrix[i][j]
			obj.sum[i][j] = matrix[i][j]
			if i-1 >= 0 {
				obj.sum[i][j] = obj.sum[i][j] + obj.sum[i-1][j]
			}

			if j-1 >= 0 {
				obj.sum[i][j] = obj.sum[i][j] + obj.sum[i][j-1]
			}

			if i-1 >= 0 && j-1 >= 0 {
				obj.sum[i][j] = obj.sum[i][j] - obj.sum[i-1][j-1]
			}
		}
	}
	//for i := 0; i < len(matrix); i++ {
	//	for j := 0; j < len(matrix[0]); j++ {
	//		fmt.Print(obj.save[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	//fmt.Println("--------------------")
	//
	//for i := 0; i < len(matrix); i++ {
	//	for j := 0; j < len(matrix[0]); j++ {
	//		fmt.Print(obj.sum[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	return *obj
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.sum[row2][col2]
	if row1-1 >= 0 {
		result = result - this.sum[row1-1][col2]
	}
	if col1-1 >= 0 {
		result = result - this.sum[row2][col1-1]
	}
	if row1-1 >= 0 && col1-1 >= 0 {
		result = result + this.sum[row1-1][col1-1]
	}
	return result
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}

	testmatrix := Constructor(matrix)
	fmt.Println(testmatrix.SumRegion(2, 1, 4, 3))

}

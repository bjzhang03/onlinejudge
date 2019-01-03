package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	result := 0
	if n > 0 {
		save := make([]int, n+1)
		save[1] = 1

		for i := 2; i <= n; i++ {
			tmpres := math.MaxInt32
			for j := 1; j < i; j++ {
				if tmpres > save[j]+save[i-j] {
					tmpres = save[j] + save[i-j]
				}
			}
			save[i] = tmpres

			if i == int(math.Sqrt(float64(i)))*int(math.Sqrt(float64(i))) {
				save[i] = 1
			}
		}
		//fmt.Println(save)
		result = save[n]
	}
	return result
}

func main() {
	fmt.Println(numSquares(16))

}

package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	result := 0
	if n > 0 {
		save := make([]int, n)
		for i := 0; i < n; i++ {
			save[i] = i
		}

		limit := int(math.Sqrt(float64(n)))
		for i := 2; i <= limit; i++ {
			if save[i] > 1 {
				for j := i + i; j < n; {
					save[j] = 1
					j = j + i
				}
			}
		}

		for i := 0; i < n; i++ {
			if save[i] > 1 {
				result++
			}
		}
	}
	return result
}

func main() {

	fmt.Println(countPrimes(1500000))

}

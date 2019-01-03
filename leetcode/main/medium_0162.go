package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	result := 0
	if len(nums) > 0 {
		pnums := []int{math.MinInt64} // processed nums save
		pnums = append(pnums, nums...)
		pnums = append(pnums, math.MinInt64)

		for i := 1; i <= len(nums); i++ {
			if pnums[i] > pnums[i-1] && pnums[i] > pnums[i+1] {
				result = i - 1
				break
			}
		}
	}
	return result
}

func main() {

	nums := []int{1}

	fmt.Print(findPeakElement(nums))

}

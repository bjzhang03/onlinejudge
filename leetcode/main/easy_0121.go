package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	if len(prices) > 0 {
		//tmpSave := make([][]int, len(prices))
		//for i := 0; i < len(prices); i++ {
		//	tmpSave[i] = make([]int, len(prices))
		//}
		//fmt.Println(tmpSave)
		for i := 0; i < len(prices); i++ {
			for j := i + 1; j < len(prices); j++ {
				if result < prices[j]-prices[i] {
					result = prices[j] - prices[i]
				}
			}
		}
	}
	return result

}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

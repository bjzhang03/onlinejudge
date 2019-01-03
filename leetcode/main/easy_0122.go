package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	if len(prices) > 0 {
		save := make([]int, len(prices))
		for i := 0; i < len(prices); i++ {
			sell := 0
			nosell := 0
			for j := 0; j <= i; j++ {
				if nosell < save[j] {
					nosell = save[j]
				}
				if sell < save[j]+prices[i]-prices[j] {
					sell = save[j] + prices[i] - prices[j]
				}
			}
			if nosell < sell {
				save[i] = sell
			} else {
				save[i] = nosell
			}
			if result < save[i] {
				result = save[i]
			}
		}
	}
	return result
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

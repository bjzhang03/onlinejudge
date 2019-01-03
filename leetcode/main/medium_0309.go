package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	// 如果只有一天的话
	// 则无法获利
	if len(prices) > 1 {
		maxbenefit := make([]int, len(prices))
		for i := 0; i < len(prices); i++ {
			// 在i天卖出,此时的最大的获利
			for j := i - 1; j >= 0; j-- {
				// 决定在j天买入的时间的话
				tmp := prices[i] - prices[j]
				// 之前获得的最优的结果
				if j-2 >= 0 {
					tmp = tmp + maxbenefit[j-2]
				}
				if result < tmp {
					result = tmp
				}
			}
			maxbenefit[i] = result
		}
	}
	return result
}

func main() {

	fmt.Println(maxProfit([]int{6, 1, 6, 4, 3, 0, 2}))

}

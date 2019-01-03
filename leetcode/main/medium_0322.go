package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	result := 0
	if len(coins) > 0 && amount > 0 {
		// 先对硬币数据进行排序
		save := make([]int, amount+1)
		for i := 0; i <= amount; i++ {
			current := math.MaxInt32
			for j := 0; j < len(coins); j++ {
				if i-coins[j] >= 0 && save[i-coins[j]] > 0 && current > save[i-coins[j]]+1 {
					//fmt.Println("i = ", i, i-coins[j], save[i-coins[j]])
					current = save[i-coins[j]] + 1
				}
				if coins[j] == i {
					current = 1
				}
			}
			if current < math.MaxInt32 {
				save[i] = current
			} else {
				save[i] = -1
			}
		}
		//fmt.Println(save)
		result = save[amount]
	}
	return result
}

func main() {

	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6259))

}

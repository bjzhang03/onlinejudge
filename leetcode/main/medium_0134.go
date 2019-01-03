package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	result := -1
	if len(gas) > 0 && len(cost) > 0 {
		doublegas := []int{}
		doublegas = append(doublegas, gas...)
		doublegas = append(doublegas, gas...)
		doublecost := []int{}
		doublecost = append(doublecost, cost...)
		doublecost = append(doublecost, cost...)
		for i := 0; i < len(gas); i++ {
			sum := 0
			j := i
			for ; j < i+len(gas); j++ {
				sum = sum + doublegas[j] - doublecost[j]
				if sum < 0 {
					break
				}
			}
			if j == i+len(gas) {
				result = i
				break
			}
		}
		//fmt.Println("doublegas ", doublegas)
		//fmt.Println("doublecost ", doublecost)
	}
	return result
}

func main() {

	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))

}

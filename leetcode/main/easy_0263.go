package main

import (
	"fmt"
)

func isUgly(num int) bool {
	result := false
	if num > 0 {
		for num%2 == 0 {
			num = num / 2
		}
		for num%3 == 0 {
			num = num / 3
		}
		for num%5 == 0 {
			num = num / 5
		}
		if num == 1 {
			result = true
		}
	}
	return result

}

func main() {

	fmt.Println(isUgly(14))

}

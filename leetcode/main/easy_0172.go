package main

import "fmt"

func trailingZeroes(n int) int {
	result := 0
	for n > 0 {
		result = result + n/5
		n = n / 5
	}
	return result

}

func main() {

	fmt.Println(trailingZeroes(3))

}

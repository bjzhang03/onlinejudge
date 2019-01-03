package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n > 0 {
		result := true

		for n > 0 {
			if n == 1 {
				break
			}
			if n%3 != 0 {
				result = false
				break
			}
			n = n / 3
		}
		return result
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(45))

}

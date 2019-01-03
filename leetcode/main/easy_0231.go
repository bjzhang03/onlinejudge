package main

import "fmt"

func isPowerOfTwo(n int) bool {
	result := true
	if n <= 0 {
		result = false
	} else {
		for n != 0 {
			if n&(0x1) == 1 && n > 1 {
				result = false
				break
			}
			n = n >> 1
		}
	}
	return result
}

func main() {
	fmt.Println(isPowerOfTwo(-4))
}

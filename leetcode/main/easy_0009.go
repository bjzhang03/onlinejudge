package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	result := false
	if x >= 0 {
		result = true
		str := strconv.Itoa(x)

		for i := 0; i <= len(str)-1-i; i++ {
			if str[i] != str[len(str)-1-i] {
				result = false
				break
			}
		}
	}
	return result

}
func main() {
	fmt.Println(isPalindrome(121))

}

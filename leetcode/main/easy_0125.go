package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	result := true
	strLower := strings.ToLower(s)
	str := ""
	for i := 0; i < len(strLower); i++ {
		if strLower[i] >= 'a' && strLower[i] <= 'z' || strLower[i] >= '0' && strLower[i] <= '9' {
			str = str + string(strLower[i])
		}
	}
	start := 0
	end := len(str) - 1
	for start <= end {
		if str[start] != str[end] {
			result = false
			break
		} else {
			start++
			end--
		}
	}
	return result
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))

}

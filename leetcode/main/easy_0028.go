package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	result := -1
	if len(haystack) > 0 && len(needle) > 0 && len(haystack) >= len(needle) {
		for i := 0; i <= len(haystack)-len(needle); i++ {
			flag := true
			for j := 0; j < len(needle); j++ {
				if i+j < len(haystack) && haystack[i+j] != needle[j] {
					flag = false
				}
			}
			if flag {
				result = i
				break
			}
		}
	} else if len(needle) == 0 {
		result = 0
	}
	return result
}
func main() {
	fmt.Println(strStr("mississippi", "issipi"))
	fmt.Println(strStr("a", "a"))
}

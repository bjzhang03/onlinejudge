package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	result := 0
	if len(s) > 0 {
		start := len(s)
		end := len(s)
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != ' ' {
				start = i
				end = i
				break
			}
		}
		for i := start; i >= 0; i-- {
			if s[i] != ' ' {
				end = i
			} else {
				break
			}
		}
		result = start - end + 1
	}
	return result
}
func main() {
	s := "   day"
	fmt.Println(lengthOfLastWord(s))

}

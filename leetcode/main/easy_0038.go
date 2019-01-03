package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	result := ""
	if n > 0 {
		result = "1"
		for i := 2; i <= n; i++ {
			tmp := ""
			for j := 0; j < len(result); j++ {
				count := 1
				for ; j+1 < len(result) && result[j] == result[j+1]; {
					count++
					j++
				}
				tmp = tmp + strconv.Itoa(count) + string(result[j])
			}
			result = tmp
		}
	}
	return result
}

func main() {
	fmt.Println(countAndSay(6))

}

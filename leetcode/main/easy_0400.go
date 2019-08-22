package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	if (n > 0) {
		result := 0;
		count := 1;
		for n > 0 {
			countStr := strconv.Itoa(count)
			if len(countStr) >= n {
				result, _ = strconv.Atoi(string(countStr[n-1]))
			}
			n = n - len(countStr)
			count++
		}
		return result
	}
	return -1;
}

func main() {

	fmt.Println(findNthDigit(1111111111))

}

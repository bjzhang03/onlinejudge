package main

import (
	"fmt"
	"math"
)

func titleToNumber(s string) int {
	result := 0
	lton := make(map[uint8]int)
	for i := 0; i < 26; i++ {
		lton[uint8('A'+i)] = i + 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		result = result + int(math.Pow(26, float64(len(s)-1-i)))*lton[s[i]]
	}
	fmt.Println(lton)
	return result
}

func main() {

	fmt.Println(titleToNumber("ZY"))

}

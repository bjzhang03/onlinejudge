package main

import "fmt"

// 定义全局变量并进行初始化的操作
var save = map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
var table = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	result := ""
	if num > 0 {
		// fmt.Println(save)
		// fmt.Println(table)
		if num >= 1000 {
			result = result + "M"
			result = result + intToRoman(num-1000)
		} else {
			for i := 1; i < len(table); i++ {
				// fmt.Println("i = ", i)
				if num >= table[i] && num < table[i-1] {
					result = result + save[table[i]]
					result = result + intToRoman(num-table[i])
					break
				}
			}

		}

	}
	return result
}
func main() {
	fmt.Println("result := ", intToRoman(1994))
}

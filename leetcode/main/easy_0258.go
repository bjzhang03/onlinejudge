package main

import "fmt"

func addDigits(num int) int {
	result := 0
	if num > 0 {
		for num >= 10 {
			tmp := 0
			for num > 0 {
				tmp = tmp + num%10
				num = num / 10
			}
			num = tmp
		}
		result = num
	}
	return result
}

func main() {

	fmt.Println(addDigits(38))

}

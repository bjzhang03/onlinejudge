package main

import "fmt"

func isHappy(n int) bool {
	result := true
	save := make(map[int]bool)
	for true {
		tmpres := 0
		for n > 0 {
			tmpres = tmpres + (n%10)*(n%10)
			n = n / 10
		}
		if tmpres == 1 {
			break
		} else {
			if _, ok := save[tmpres]; ok {
				result = false
				break
			} else {
				save[tmpres] = true
				n = tmpres
			}
		}

	}
	return result
}

func main() {

	fmt.Println(isHappy(2))

}

package main

import "fmt"

func convertToTitle(n int) string {
	result := ""
	if n > 0 {
		save := make(map[int]uint8)
		for i := 1; i <= 26; i++ {
			save[i] = uint8('A' + i - 1)
		}
		if n%26 == 0 {
			tmp := (n - 26) / 26
			result = result + convertToTitle(tmp) + string('Z')
		} else {
			reminder := n % 26
			tmp := n / 26
			result = result + convertToTitle(tmp) + string(save[reminder])
		}
	}
	return result
}

func main() {
	fmt.Println(convertToTitle(52))

}

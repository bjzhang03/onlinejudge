package main

import "fmt"

func romanToInt(s string) int {
	result := 0
	m := make(map[uint8]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	//fmt.Println(m['I'])
	for i := 0; i < len(s); i++ {
		//fmt.Println(i, " - ", result)
		if i+1 <= len(s)-1 && m[s[i]] < m[s[i+1]] {
			result = result - m[s[i]]
		} else {
			result = result + m[s[i]]
		}
	}
	return result
}
func main() {
	fmt.Println("result = ", romanToInt("IV"))
}

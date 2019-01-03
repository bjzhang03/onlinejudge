package main

import "fmt"

var table = map[string]uint8{"1": 'A', "2": 'B', "3": 'C', "4": 'D', "5": 'E', "6": 'F', "7": 'G', "8": 'H', "9": 'I', "10": 'J', "11": 'K', "12": 'L', "13": 'M',
	"14": 'N', "15": 'O', "16": 'P', "17": 'Q', "18": 'R', "19": 'S', "20": 'T', "21": 'U', "22": 'V', "23": 'W', "24": 'X', "25": 'Y', "26": 'Z'}

func numDecodings(s string) int {
	result := 0
	if len(s) == 1 {
		if _, ok := table[string(s[0])]; ok {
			result = 1
		}
	} else if len(s) == 2 {
		if _, ok := table[s]; ok {
			result = result + 1
		}

		if _, ok := table[string(s[0])]; ok {
			if _, ok2 := table[string(s[1])]; ok2 {
				result = result + 1
			}
		}
	} else if len(s) > 2 {
		save := make([]int, len(s))
		if _, ok := table[string(s[0])]; ok {
			save[0] = 1
		}
		if _, ok := table[string(s[0])+string(s[1])]; ok {
			save[1] = save[1] + 1
		}
		if _, ok := table[string(s[0])]; ok {
			if _, ok2 := table[string(s[1])]; ok2 {
				save[1] = save[1] + 1
			}
		}
		for i := 2; i < len(s); i++ {
			if _, ok := table[string(s[i])]; ok {
				save[i] = save[i] + save[i-1]
				//fmt.Println("save = ", save[i])
			}
			if _, ok := table[string(s[i-1])+string(s[i])]; ok {
				save[i] = save[i] + save[i-2]
				//fmt.Println("save = ", save[i])
			}
		}
		//fmt.Println(save)
		result = save[len(s)-1]
	}
	return result
}

func main() {
	fmt.Println(numDecodings("226"))
}

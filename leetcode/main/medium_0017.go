package main

import "fmt"

var save = map[uint8]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

func dfs(digits string, result *[]string, index int, tmpsave string) {
	// fmt.Println("tmpsave = ", tmpsave)
	// 边界条件判断
	if len(tmpsave) == len(digits) {
		*result = append(*result, tmpsave)
		// fmt.Println(*result)
		return
	}

	for i := index; i < len(digits); i++ {
		for j := 0; j < len(save[digits[i]]); j++ {
			tmp := tmpsave + string(save[digits[i]][j])
			dfs(digits, result, i+1, tmp)
		}
	}
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) > 0 {
		tmpsave := ""
		dfs(digits, &result, 0, tmpsave)
	}
	return result
}

func main() {

	fmt.Println(letterCombinations("23"))

}

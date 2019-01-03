package main

import "fmt"

// 判断是否是回文
func isPalindrome(s string) bool {
	result := true
	if len(s) > 0 {
		ins := 0
		ine := len(s) - 1
		for ins <= ine {
			if s[ins] != s[ine] {
				result = false
				break
			}
			ins++
			ine--
		}
	}
	return result
}

func dfs(result *[][]string, s string, index int, save []string) {
	//fmt.Println(index, used)
	if index == len(s) {
		tmpsave := []string{}
		tmpsave = append(tmpsave, save...)
		*result = append(*result, tmpsave)
		//fmt.Println(save)
	} else if index < len(s) {
		tmp := ""
		for j := index; j < len(s); j++ {
			tmp = tmp + string(s[j])
			if isPalindrome(tmp) {
				save = append(save, tmp)
				dfs(result, s, j+1, save)
				save = save[:len(save)-1]
			}
		}
	}
}

func partition(s string) [][]string {
	result := [][]string{}
	if len(s) > 0 {
		dfs(&result, s, 0, []string{})
	}
	return result
}

func main() {

	fmt.Println(partition("seeslaveidemonstrateyetartsnomedievalsees"))

}

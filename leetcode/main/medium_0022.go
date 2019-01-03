package main

import "fmt"

func hasstr(save []string, str string) bool {
	result := false
	for i := 0; i < len(save); i++ {
		if str == save[i] {
			result = true
			break
		}
	}
	return result
}

func leftCount(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			count++
		}
	}
	return count
}
func rightCount(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ')' {
			count++
		}
	}
	return count
}
func dfs(result *[]string, tmp string, n int) {
	if len(tmp) == 2*n {
		if !hasstr(*result, tmp) {
			*result = append(*result, tmp)
		}
		return
	}
	if leftCount(tmp) < n {
		nexttmp := tmp + string('(')
		dfs(result, nexttmp, n)
	}
	if leftCount(tmp) == n && rightCount(tmp) < n {
		nexttmp := tmp + string(')')
		dfs(result, nexttmp, n)
	}
	if rightCount(tmp) < leftCount(tmp) {
		nexttmp := tmp + string(')')
		dfs(result, nexttmp, n)
	}

}
func generateParenthesis(n int) []string {
	if n > 0 {
		result := []string{}
		tmp := ""
		dfs(&result, tmp, n)
		return result
	}
	return []string{}
}

func main() {
	fmt.Println(generateParenthesis(3))

}

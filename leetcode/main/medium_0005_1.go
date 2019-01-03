package main

import "fmt"

func longestPalindrome(s string) string {
	result := ""
	if len(s) > 0 {
		// 初始化dp的数据全部为false
		dpsave := make([][]bool, len(s))
		for i := 0; i < len(s); i++ {
			dpsave[i] = make([]bool, len(s))
		}

		for slen := 0; slen < len(s); slen++ {
			for i := 0; i < len(s)-slen; i++ {
				if slen == 0 {
					// 处理长度为1的情况
					dpsave[i][i+slen] = true
				} else if slen == 1 && s[i] == s[i+slen] {
					// 处理偶数情况
					dpsave[i][i+slen] = true
				} else if slen != 1 && s[i] == s[i+slen] {
					dpsave[i][i+slen] = dpsave[i+1][i+slen-1]
				}

				// 如果当前的回文比以前的长，则更新以前的回文数据
				if len(result) < slen+1 && dpsave[i][i+slen] {
					result = ""
					for t := i; t <= i+slen; t++ {
						result = result + string(s[t])
					}

				}
			}
		}
		fmt.Println(dpsave)
	}
	return result
}

func main() {

	fmt.Println(longestPalindrome("babad"))

}

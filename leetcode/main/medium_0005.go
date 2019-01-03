package main

import "fmt"

// https://blog.csdn.net/sodaoo/article/details/78564053
func longestPalindrome(s string) string {
	result := ""
	//把字符串变成长度为奇数的数据
	substr := "#"
	for i := 0; i < len(s); i++ {
		substr = substr + string(s[i])
		substr = substr + "#"
	}
	// 以每一个字符为中心向两边进行判断操作
	for i := 0; i < len(substr); i++ {
		length := 0
		for ; i-length >= 0 && i+length < len(substr); length++ {
			if substr[i-length] != substr[i+length] {
				break
			}
		}
		// 取出当前的数据来进行处理
		if length >= len(result) {
			tmp := ""
			for j := 0; j < length; j++ {
				if substr[i-j] != '#' {
					tmp = string(substr[i-j]) + tmp
				}

				if i-j != i+j && substr[i+j] != '#' {
					tmp = tmp + string(substr[i+j])
				}
			}
			// 认为第一个数据是最长的
			if len(result) < len(tmp) {
				result = tmp
			}
		}
	}

	return result
}
func main() {
	fmt.Println(longestPalindrome("babad"))
}

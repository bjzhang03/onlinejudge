package main

func isPrefix(prefix string, str string) bool {
	result := true
	for i := 0; i < len(prefix); i++ {
		if str[i] != prefix[i] {
			result = false
			break
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) > 0 {
		// 找出长度最短的str
		result = strs[0]
		for i := 0; i < len(strs); i++ {
			if len(result) > len(strs[i]) {
				result = strs[i]
			}
		}

		for i := len(result); i >= 0; i-- {
			tmp := result[0:i]
			flag := true
			for j := 0; j < len(strs); j++ {
				if !isPrefix(tmp, strs[j]) {
					flag = false
					break
				}
			}
			if flag {
				result = tmp
				break
			}
		}
	}
	return result
}
func main() {


}

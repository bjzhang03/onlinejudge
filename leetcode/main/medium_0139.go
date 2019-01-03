package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	result := false
	if len(s) == 0 {
		result = true
	} else if len(s) > 0 && len(wordDict) > 0 {
		save := make([]bool, len(s))
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(wordDict); j++ {
				if i+1 == len(wordDict[j]) {
					//fmt.Println("=", s[i+1-len(wordDict[j]):i+1])
					if wordDict[j] == s[i+1-len(wordDict[j]):i+1] {
						save[i] = true
					}

				} else if i+1 > len(wordDict[j]) {
					//fmt.Println(">", s[i+1-len(wordDict[j]):i+1])
					if wordDict[j] == s[i+1-len(wordDict[j]):i+1] && save[i+1-len(wordDict[j])-1] {
						save[i] = true
					}

				}
			}
		}
		//fmt.Println(save)
		result = save[len(save)-1]
	}
	return result
}

func main() {
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak("applepenapple", wordDict))

}

package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	result := []string{}
	if len(s) > 0 {
		used := make(map[string]bool)
		apperence := make(map[string]bool)
		for i := 0; i <= len(s)-10; i++ {
			tmp := s[i : i+10]
			if _, ok := apperence[tmp]; ok {
				if _, uok := used[tmp]; !uok {
					result = append(result, tmp)
					used[tmp] = true
				}
			} else {
				apperence[tmp] = true
			}
		}
	}
	return result
}

func main() {

	fmt.Print(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))

}

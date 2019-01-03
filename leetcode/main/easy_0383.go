package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	result := false
	if len(magazine) >= len(ransomNote) {
		result = true
		countRansomNote := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		countMagazine := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		for i := 0; i < len(ransomNote); i++ {
			countRansomNote[ransomNote[i]]++
		}
		for i := 0; i < len(magazine); i++ {
			countMagazine[magazine[i]]++
		}
		for k, v := range countRansomNote {
			if v > countMagazine[k] {
				result = false
			}
		}
	}
	return result
}

func main() {
	fmt.Println(canConstruct("aaa", "aab"))
}

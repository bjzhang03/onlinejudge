package main

import "fmt"

func reverseVowels(s string) string {
	if len(s) > 0 {
		// 构造set的数据集合
		vowels := make(map[uint8]bool)
		vowels['a'] = true
		vowels['o'] = true
		vowels['e'] = true
		vowels['i'] = true
		vowels['u'] = true
		vowels['A'] = true
		vowels['O'] = true
		vowels['E'] = true
		vowels['I'] = true
		vowels['U'] = true
		// 将string 转化成slice
		result := []uint8{}
		for _, k := range s {
			result = append(result, uint8(k))
		}
		// 双指针查找数据
		start := 0
		end := len(result) - 1
		for start < end {
			for start < len(result) {
				if !vowels[result[start]] {
					start++
				} else {
					break
				}
			}
			for end >= 0 {
				if !vowels[result[end]] {
					end--
				} else {
					break
				}
			}
			if start < end {
				tmp := result[start]
				result[start] = result[end]
				result[end] = tmp
				start++
				end--
			}
		}
		// 将数据转化成结果
		str := ""
		for _, k := range result {
			str = str + string(k)
		}
		return str
	}
	return ""
}

func main() {

	fmt.Println(reverseVowels("OE"))

}

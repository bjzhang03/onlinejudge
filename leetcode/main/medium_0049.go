package main

import (
	"fmt"
)

func mapequal(a map[uint8]int, b map[uint8]int) bool {
	result := true
	if len(a) != len(b) {
		result = false
	} else {
		for k, v := range a {
			if _, ok := b[k]; !ok {
				// k不存在
				result = false
				break
			} else if b[k] != v {
				// 存在但是不相等
				result = false
				break
			}
		}
	}
	return result
}

func strtomap(str string) map[uint8]int {
	result := make(map[uint8]int)
	if len(str) > 0 {
		for i := 0; i < len(str); i++ {
			if _, ok := result[str[i]]; !ok {
				result[str[i]] = 1
			} else {
				result[str[i]]++
			}
		}
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	fmt.Println(len(strs))
	if len(strs) > 0 {
		// 构建下一次递归的数据集合
		usedindex := make(map[int]bool)
		for i := 0; i < len(strs); i++ {
			// 如果当前的数据是没有被使用过的数据
			if _, ok := usedindex[i]; !ok {
				fmt.Println("i = ", i)
				save := []string{strs[i]}
				savemap := strtomap(strs[i])
				for j := i + 1; j < len(strs); j++ {
					if mapequal(savemap, strtomap(strs[j])) {
						fmt.Println("j = ", j)
						save = append(save, strs[j])
						usedindex[j] = true
					}
				}
				usedindex[i] = true
				result = append(result, save)
			}
			// fmt.Println(strtomap(strs[i]))
		}
	}
	return result
}
func main() {
	a := map[uint8]int{'a': 1}
	b := map[uint8]int{'a': 1}
	if mapequal(a, b) {
		fmt.Println("aaaa")
	}
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}

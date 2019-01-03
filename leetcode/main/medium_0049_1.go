package main

import (
	"fmt"
	"sort"
)

func sortstring(s string) string {
	result := ""
	save := []int{}
	for i := 0; i < len(s); i++ {
		save = append(save, int(s[i]))
	}
	sort.Ints(save)
	for i := 0; i < len(save); i++ {
		result = result + string(uint8(save[i]))
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	// fmt.Println(len(strs))
	if len(strs) > 0 {
		// 构建下一次递归的数据集合
		savemap := make(map[string][]string)
		usedmap := make(map[string]bool)
		for i := 0; i < len(strs); i++ {
			// fmt.Println(sortstring(strs[i]))
			tmpstr := sortstring(strs[i])
			if _, ok := usedmap[tmpstr]; !ok {
				savemap[tmpstr] = []string{strs[i]}
				usedmap[tmpstr] = true
			} else {
				savemap[tmpstr] = append(savemap[tmpstr], strs[i])
			}
			// fmt.Println(strtomap(strs[i]))
		}
		for _, v := range savemap {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(sortstring("tea"))

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, " ")
	nums := strings.Split(lineArr[0], ",")
	target, _ := strconv.Atoi(lineArr[1])
	//fmt.Println("target = ", target)
	savemap := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := savemap[nums[i]]; !ok {
			savemap[nums[i]] = 1
		} else {
			savemap[nums[i]]++
		}
	}
	countSave := [][]int{}
	for key, val := range savemap {
		tmpkey, _ := strconv.Atoi(key)
		countSave = append(countSave, []int{tmpkey, val})
	}
	sort.Slice(countSave, func(i, j int) bool {
		return comparsLess(countSave[i], countSave[j])
	})
	sta := 0
	result := ""
	for ; sta < target-1; sta++ {
		result = result + strconv.Itoa(countSave[sta][0]) + ","
	}
	result = result + strconv.Itoa(countSave[sta][0])
	//fmt.Println(countSave)
	// 返回处理后的结果
	// return ans
	return result
}

func comparsLess(a []int, b []int) bool {
	if a[1] > b[1] {
		return true
	}else if a[1]==b[1]{
		return a[0]<b[0]
	}
	return false
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

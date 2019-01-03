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
	lineArr := strings.Split(line, ",")
	nums := []int{}
	for i := 0; i < len(lineArr); i++ {
		tmp, _ := strconv.Atoi(lineArr[i])
		nums = append(nums, tmp)
	}
	sort.Ints(nums)
	// 返回处理后的结果
	// return ans
	return strconv.Itoa(nums[len(nums)/2])
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

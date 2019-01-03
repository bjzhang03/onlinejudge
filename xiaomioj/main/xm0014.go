package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, " ")
	numsArr := lineArr[0]
	nums := strings.Split(numsArr, ",")
	target, _ := strconv.Atoi(lineArr[1])

	for i := 0; i < len(nums); i++ {
		tmp, _ := strconv.Atoi(nums[i])
		if tmp == target {
			return strconv.Itoa(i)
		}
	}
	// 返回处理后的结果
	// return ans
	return "-1"
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

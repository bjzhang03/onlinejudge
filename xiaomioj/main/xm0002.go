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
	result := 0
	for i := 0; i < len(lineArr); i++ {
		tmpNum, _ := strconv.Atoi(lineArr[i])
		result = result ^ tmpNum
	}
	// 返回处理后的结果
	// return ans
	return strconv.Itoa(result)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, " ")
	strA := lineArr[0]
	strB := lineArr[1]
	countA := make(map[uint8]int)
	countB := make(map[uint8]int)
	for i := 0; i < len(strA); i++ {
		if val, ok := countA[strA[i]]; ok {
			countA[strA[i]] = val + 1
		} else {
			countA[strA[i]] = 1
		}
	}
	for i := 0; i < len(strB); i++ {
		if val, ok := countB[strB[i]]; ok {
			countB[strB[i]] = val + 1
		} else {
			countB[strB[i]] = 1
		}
	}
	for val, _ := range countA {
		if countA[val] > countB[val] {
			return "false"
		}
	}
	// 返回处理后的结果
	// return ans
	return "true"
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

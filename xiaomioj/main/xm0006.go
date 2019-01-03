package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, ",")
	strA := lineArr[0]
	strB := lineArr[1]
	target := lineArr[2]
	if len(strA)+len(strB) == len(target) {
		return crossQueue(strA, strB, target)
	}
	// 返回处理后的结果
	// return ans
	return "false"
}

func crossQueue(strA string, strB string, target string) string {
	// 初始化数据
	content := make([][][]string, len(strA)+1)
	for i := 0; i < len(strA)+1; i++ {
		content[i] = make([][]string, len(strB)+1)
	}
	content[0][0] = []string{""}
	for i := 1; i < len(strA)+1; i++ {
		content[i][0] = []string{strA[:i]}
	}
	for i := 1; i < len(strB)+1; i++ {
		content[0][i] = []string{strB[:i]}
	}
	for i := 1; i < len(strA)+1; i++ {
		for j := 1; j < len(strB)+1; j++ {
			tmp := []string{}
			for k := 0; k < len(content[i][j-1]); k++ {
				tmp = append(tmp, content[i][j-1][k]+string(strB[j-1]))
			}
			for k := 0; k < len(content[i-1][j]); k++ {
				tmp = append(tmp, content[i-1][j][k]+string(strA[i-1]))
			}
			content[i][j] = tmp
		}
	}
	for i := 0; i < len(content[len(strA)][len(strB)]); i++ {
		if content[len(strA)][len(strB)][i] == target {
			return "true"
		}
	}
	return "false"
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

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
	lineArr := strings.Split(line, ",")
	save := make(map[int]bool)
	// 将数据添加到map中去
	for i := 0; i < len(lineArr); i++ {
		tmp, _ := strconv.Atoi(lineArr[i])
		if _, ok := save[tmp]; !ok {
			save[tmp] = true
		}
	}

	result := 1
	for key, _ := range save {
		staKey := key
		// 向前找
		for save[staKey] {
			if _, ok := save[staKey-1]; !ok {
				break
			}
			staKey = staKey - 1
		}
		// 向后找
		endKey := key

		for save[endKey] {
			if _, ok := save[endKey+1]; !ok {
				break
			}
			endKey = endKey + 1
		}
		if result < endKey-staKey+1 {
			result = endKey - staKey + 1
		}
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

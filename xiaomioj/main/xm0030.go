package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(line string) string {
	// please write your code here
	num, _ := strconv.Atoi(line)

	numStr := strconv.FormatInt(int64(num), 2)
	for len(numStr) < 32 {
		numStr = "0" + numStr
	}
	//fmt.Println(numStr)

	resStr := ""
	for i := len(numStr) - 1; i >= 0; i-- {
		resStr = resStr + string(numStr[i])
	}
	result, _ := strconv.ParseInt(resStr, 2, 64)

	// 返回处理后的结果
	// return ans
	return fmt.Sprint(uint32(result))
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(line string) string {
	// please write your code here
	count, _ := strconv.Atoi(line)
	save := make([]int, count+1)
	save[1] = 1
	save[2] = 2
	for i := 3; i <= count; i++ {
		save[i] = save[i-1] + save[i-2]
	}
	// 返回处理后的结果
	// return ans
	return strconv.Itoa(save[count])
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

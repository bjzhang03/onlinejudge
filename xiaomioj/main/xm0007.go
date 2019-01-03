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
	save := make([]int, len(lineArr)+10)
	//fmt.Println("save = ", save)
	for i := 0; i < len(lineArr); i++ {
		tmp, _ := strconv.Atoi(lineArr[i])
		if tmp > 0 && tmp < len(save) {
			save[tmp] = tmp
		}
	}

	//fmt.Println("save = ", save)
	for i := 1; i < len(save); i++ {
		if save[i] != i {
			return strconv.Itoa(i)
		}
	}

	// 返回处理后的结果
	// return ans
	return strconv.Itoa(len(save) + 1)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

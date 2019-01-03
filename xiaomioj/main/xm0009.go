package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, " ")
	nums := lineArr[0]
	count, _ := strconv.Atoi(lineArr[1])
	result := math.MaxInt64
	deepfirsts(nums, len(nums)-count, &result, 0, "")
	// 返回处理后的结果
	// return ans
	return strconv.Itoa(result)
}

func deepfirsts(nums string, count int, result *int, index int, tmps string) {
	if len(tmps) == count {
		tmpn, _ := strconv.Atoi(tmps)
		if *result > tmpn {
			*result = tmpn
		}
		//fmt.Println("result = ", *result)
	} else if index < len(nums) {
		for i := index; i < len(nums); i++ {
			tmps = tmps + string(nums[i])
			deepfirsts(nums, count, result, i+1, tmps)
			tmps = tmps[:len(tmps)-1]
		}
	}
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

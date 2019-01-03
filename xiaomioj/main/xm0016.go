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
	result := "err"
	// 对输入的数据进行切分操作
	lineArr := strings.Split(line, " ")
	//fmt.Println("lineArry = ", lineArr, len(lineArr))
	operators := []string{}
	nums := []int{}
	// 处理乘法与出发的过程
	for i := 0; i < len(lineArr); i++ {
		// 对操作符的处理
		if lineArr[i] == "+" || lineArr[i] == "-" || lineArr[i] == "*" || lineArr[i] == "/" {
			operators = append(operators, lineArr[i])
		} else {
			tmp, _ := strconv.Atoi(lineArr[i])
			if len(operators) > 0 && operators[len(operators)-1] == "/" {
				if tmp == 0 {
					return result
				} else {
					tmp = nums[len(nums)-1] / tmp
					nums[len(nums)-1] = tmp
				}
				operators = operators[:len(operators)-1]

			} else if len(operators) > 0 && operators[len(operators)-1] == "*" {
				tmp = nums[len(nums)-1] * tmp
				nums[len(nums)-1] = tmp
				operators = operators[:len(operators)-1]
			} else {
				nums = append(nums, tmp)
			}
		}
	}
	//fmt.Println(operators, nums)
	// 对加法与减法的处理

	numRes := nums[0]
	for i := 1; i < len(nums); i++ {
		if operators[i-1] == "+" {
			numRes = numRes + nums[i]
		} else if operators[i-1] == "-" {
			numRes = numRes - nums[i]
		}
	}
	result = strconv.Itoa(numRes)

	// 返回处理后的结果
	// return ans
	return result
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

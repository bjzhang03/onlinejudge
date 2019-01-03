package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var flag = 1
	// 判断数字的符号
	if x < 0 {
		flag = -1
	}
	// 设置结果变量
	result := 0
	x = int(math.Abs(float64(x)))
	// 每次取的都是数据最后的部分的数据
	for x > 0 {
		tmp := x % 10
		result = result*10 + tmp
		x = x / 10
	}
	// 判断数据是不是已经超出范围了
	if result >= math.MaxInt32 {
		result = 0
	}
	return result * flag
}

func main() {
	fmt.Print(reverse(123))

}

package main

import (
	"fmt"
	"strconv"
)

func isOperator(operator uint8) int {
	result := 0
	if operator == '+' {
		result = 1
	} else if operator == '-' {
		result = 2
	} else if operator == '*' {
		result = 3
	} else if operator == '/' {
		result = 4
	}
	return result
}
func handleMul(nums []int, opes []uint8) ([]int, []uint8) {
	resnums := []int{nums[0]}
	resopes := []uint8{}
	for i := 1; i < len(nums); i++ {
		//fmt.Println(resnums, " - ", resopes)
		if opes[i-1] == '*' {
			resnums[len(resnums)-1] = resnums[len(resnums)-1] * nums[i]
		} else if opes[i-1] == '/' {
			resnums[len(resnums)-1] = resnums[len(resnums)-1] / nums[i]
		} else {
			resnums = append(resnums, nums[i])
			resopes = append(resopes, opes[i-1])
		}
	}
	return resnums, resopes
}

func handleAdd(nums []int, opes []uint8) int {
	result := nums[0]
	for i := 0; i < len(opes); i++ {
		if opes[i] == '+' {
			result = result + nums[i+1]
		} else if opes[i] == '-' {
			result = result - nums[i+1]
		}
	}
	return result
}
func calculate(s string) int {
	result := 0
	if len(s) > 0 {
		nums := []int{}
		opes := []uint8{} // operators

		curs := ""
		for i := 0; i < len(s); i++ {
			if isOperator(s[i]) > 0 {
				// 当前位置是操作符
				tmpnum, _ := strconv.Atoi(curs)
				// 判断前面的数字是不是乘法或者除法
				nums = append(nums, tmpnum)
				opes = append(opes, s[i])
				// 清空当前存储的数据
				curs = ""
			} else if s[i] != ' ' {
				// 当前位置是数字
				curs = curs + string(s[i])
			} else if s[i] == ' ' {
				// 当前位置是空格
			}
		}
		// 处理最后一个数据
		if len(curs) > 0 {
			tmpnum, _ := strconv.Atoi(curs)
			nums = append(nums, tmpnum)
		}
		//fmt.Println(nums, opes)
		//nums, opes = handleMul(nums, opes)
		//fmt.Println("nums = ", nums, "opes = ", opes)
		//result = handleAdd(nums, opes)
		result = handleAdd(handleMul(nums, opes))
	}
	return result
}

func main() {

	fmt.Println("result = ", calculate(" 1+5 / 2 *1 9 "))
	fmt.Println('+', '-', '*', '/')

}

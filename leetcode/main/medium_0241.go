package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isOperator(ope uint8) bool {
	result := false
	if ope == '*' || ope == '+' || ope == '-' {
		result = true
	}
	return result
}

func handler(nums []int, opes []uint8) []int {
	result := []int{}
	if len(opes) > 0 {
		for i := 0; i < len(opes); i++ {
			leftnums := []int{}
			leftopes := []uint8{}
			rightnums := []int{}
			rightopes := []uint8{}

			leftnums = append(leftnums, nums[:i+1]...)
			leftopes = append(leftopes, opes[:i]...)
			leftres := handler(leftnums, leftopes)

			rightnums = append(rightnums, nums[i+1:]...)
			rightopes = append(rightopes, opes[i+1:]...)
			rightres := handler(rightnums, rightopes)

			if opes[i] == '*' {
				for j := 0; j < len(leftres); j++ {
					for k := 0; k < len(rightres); k++ {
						result = append(result, leftres[j]*rightres[k])
					}
				}
			} else if opes[i] == '+' {
				for j := 0; j < len(leftres); j++ {
					for k := 0; k < len(rightres); k++ {
						result = append(result, leftres[j]+rightres[k])
					}
				}
			} else if opes[i] == '-' {
				for j := 0; j < len(leftres); j++ {
					for k := 0; k < len(rightres); k++ {
						result = append(result, leftres[j]-rightres[k])
					}
				}
			}
		}
	} else if len(opes) == 0 {
		result = append(result, nums[0])
	}
	//fmt.Println("nums = ", nums, "opes = ", opes, "result = ", result)
	return result
}
func diffWaysToCompute(input string) []int {
	if len(input) > 0 {
		nums := []int{}
		opes := []uint8{} // operators

		curs := ""
		for i := 0; i < len(input); i++ {
			if isOperator(input[i]) {
				// 当前位置是操作符
				tmpnum, _ := strconv.Atoi(curs)
				// 判断前面的数字是不是乘法或者除法
				nums = append(nums, tmpnum)
				opes = append(opes, input[i])
				// 清空当前存储的数据
				curs = ""
			} else if input[i] != ' ' {
				// 当前位置是数字
				curs = curs + string(input[i])
			} else if input[i] == ' ' {
				// 当前位置是空格
			}
		}
		// 处理最后一个数据
		if len(curs) > 0 {
			tmpnum, _ := strconv.Atoi(curs)
			nums = append(nums, tmpnum)
		}

		//fmt.Println(nums, opes)
		tmpres := handler(nums, opes)
		sort.Ints(tmpres)
		return tmpres
	}

	return []int{}
}

func main() {
	input := "2*3-4*5"
	fmt.Println(diffWaysToCompute(input))

}

package main

import (
	"fmt"
	"strconv"
)

// 时间实在是太长了
func validateCorrect(save []int) bool {
	result := false
	if len(save) >= 3 {
		result = true
		for i := len(save) - 1; i >= 2; i-- {
			if save[i] != save[i-1]+save[i-2] {
				result = false
				break
			}
		}
	}
	return result
}

func dfs(num string, save []int, ind int) bool {
	if ind == len(num) {
		//fmt.Println(save)
		return validateCorrect(save)
	} else if ind < len(num) {
		tmp := ""
		for i := ind; i < len(num); i++ {
			tmp = tmp + string(num[i])
			tmpnum, _ := strconv.Atoi(tmp)
			save = append(save, tmpnum)
			if dfs(num, save, i+1) {
				return true
			}
			save = save[:len(save)-1]
			// 处理不能出现前导0的情况，如果出现前导0
			// 则不再进行添加操作
			if num[i] == '0' && tmpnum == 0 {
				break
			}
		}
	}
	return false
}

func isAdditiveNumber(num string) bool {
	result := false
	if len(num) > 0 {
		result = dfs(num, []int{}, 0)
	}
	return result
}

func main() {
	fmt.Println("result = ", isAdditiveNumber("221474836472147483649"))
}

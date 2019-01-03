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
	lineArr := strings.Split(line, "-")
	numA := lineArr[0]
	numB := lineArr[1]
	// 返回处理后的结果
	// return ans
	result := subtraction(numA, numB)
	ista := 0
	for ; ista < len(result); {
		if result[ista] == '0' {
			ista++
		} else {
			break
		}
	}
	//fmt.Println("ista = ", ista)
	result = result[ista:]
	return result
}

// 计算相减的结果
func subtraction(numa string, numb string) string {
	result := ""
	aEnd := len(numa) - 1
	bEnd := len(numb) - 1

	if aEnd >= bEnd {
		carry := 0
		for aEnd >= 0 && bEnd >= 0 {
			aNum, _ := strconv.Atoi(string(numa[aEnd]))
			bNum, _ := strconv.Atoi(string(numb[bEnd]))
			nextCarry := 0
			if aNum+carry < bNum {
				nextCarry = -1
				aNum = aNum + 10
			} else {
				nextCarry = 0
			}
			result = strconv.Itoa(aNum+carry-bNum) + result
			carry = nextCarry
			aEnd--
			bEnd--
		}

		for aEnd >= 0 {
			//fmt.Println("carry = ", carry, " aEnd = ", aEnd)
			aNum, _ := strconv.Atoi(string(numa[aEnd]))
			nextCarry := 0
			if aNum+carry >= 0 {
				nextCarry = 0
				result = strconv.Itoa(aNum+carry) + result
			} else {
				nextCarry = -1
				aNum = aNum + 10
				result = strconv.Itoa(aNum+carry) + result
			}
			carry = nextCarry
			aEnd--
		}
		//fmt.Println("result = ", result)
		return result
	}
	return result
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

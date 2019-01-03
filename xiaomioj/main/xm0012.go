package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(line string) string {
	// please write your code here
	lineArr := strings.Split(line, " ")
	nums := strings.Split(lineArr[0], ",")
	target, _ := strconv.Atoi(lineArr[1])
	numsave := []int{}
	for i := 0; i < len(nums); i++ {
		tmpn, _ := strconv.Atoi(nums[i])
		numsave = append(numsave, tmpn)
	}
	sort.Ints(numsave)
	result := [][]int{}
	deepfirsts(numsave, &result, target, []int{})
	// 返回处理后的结果
	return strconv.Itoa(len(result))
}

func sumslice(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result + nums[i]
	}
	return result
}

func deepfirsts(numsave []int, result *[][]int, target int, tmpsum []int) {
	if sumslice(tmpsum) == target {
		sumdata := []int{}
		sumdata = append(sumdata, tmpsum...)
		*result = append(*result, sumdata)
	} else if sumslice(tmpsum) < target {
		for i := 0; i < len(numsave); i++ {
			tmpsum = append(tmpsum, numsave[i])
			deepfirsts(numsave, result, target, tmpsum)
			tmpsum = tmpsum[:len(tmpsum)-1]
		}
	}
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

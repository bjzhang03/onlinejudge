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
	lineArr := strings.Split(line, ",")
	// 将数据从string转化成int
	numSave := []int{}
	for i := 0; i < len(lineArr); i++ {
		tmp, _ := strconv.Atoi(lineArr[i])
		numSave = append(numSave, tmp)
	}
	sort.Ints(numSave)
	result := 0
	for i := 0; i < len(numSave)-2; i++ {
		sum := 0 - numSave[i]
		sta := i + 1
		end := len(numSave) - 1
		for sta < end {
			if numSave[sta]+numSave[end] == sum {
				result++
				// 去掉重复出现的数据
				for sta+1 <= end && numSave[sta] == numSave[sta+1] {
					sta++
				}
				sta++
			} else if numSave[sta]+numSave[end] < sum {
				sta++
			} else if numSave[sta]+numSave[end] > sum {
				end--
			}
		}
		// 去掉重复出现的数据
		for i+1 < len(numSave) && numSave[i+1] == numSave[i] {
			i++
		}
	}
	// 返回处理后的结果
	return strconv.Itoa(result)
}
func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

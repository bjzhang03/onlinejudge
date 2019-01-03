package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stu struct {
	high  int
	count int
}

func (s *stu) String() string {
	return fmt.Sprint(s.high, " ", s.count)
}

func solution(line string) string {
	// please write your code here
	result := ""
	lineArr := strings.Split(line, " ")
	count, _ := strconv.Atoi(lineArr[0])
	save := []*stu{}
	for i := 1; i < len(lineArr)-1; i = i + 2 {
		tmp := &stu{0, 0}
		tmp.high, _ = strconv.Atoi(lineArr[i])
		tmp.count, _ = strconv.Atoi(lineArr[i+1])
		save = append(save, tmp)
	}
	numRes := []*stu{}
	dfs(save, count, []*stu{}, &numRes)
	for i := 0; i < len(numRes); i++ {
		result = result + fmt.Sprint(numRes[i].high) + " " + fmt.Sprint(numRes[i].count) + " "
	}
	// 返回处理后的结果
	return result
}

func dfs(save []*stu, count int, tmpSave []*stu, result *[]*stu) bool {
	if len(tmpSave) == count {
		if stuCheck(tmpSave) {
			*result = append(*result, tmpSave[:]...)
			return true
		}
	} else if len(tmpSave) < count {
		for i := 0; i < len(save); i++ {
			tmpSave = append(tmpSave, save[i])
			nextSave := []*stu{}
			nextSave = append(nextSave, save[:i]...)
			nextSave = append(nextSave, save[i+1:]...)
			if dfs(nextSave, count, tmpSave, result) {
				return true
			}
			tmpSave = tmpSave[:len(tmpSave)-1]
		}
	}
	return false
}

// 判断当前的排列是不是正确的
func stuCheck(stuSave []*stu) bool {
	result := true

	for i := 0; i < len(stuSave); i++ {
		highcount := 0
		for j := i - 1; j >= 0; j-- {
			if stuSave[j].high >= stuSave[i].high {
				highcount++
			}
		}
		if highcount != stuSave[i].count {
			result = false
			break
		}
	}
	return result
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}

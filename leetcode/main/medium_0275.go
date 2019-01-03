package main

import "fmt"

// https://blog.csdn.net/u014673347/article/details/48226581
func hIndex(citations []int) int {
	result := 0
	if len(citations) > 0 {

		ins := 0
		ine := len(citations) - 1

		for ins <= ine {
			mid := (ins + ine) / 2
			// 对满足的结果进行一次又一次的测试
			if citations[mid] >= len(citations)-mid {
				ine = mid - 1
				result = len(citations) - mid
			} else {
				ins = mid + 1
			}
		}

	}
	return result

}

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(hIndex(nums))

}

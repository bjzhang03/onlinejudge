package main

import "fmt"

func maxInt(vals ...int) int {
	result := vals[0]
	for _, val := range vals {
		if result < val {
			result = val
		}
	}
	return result
}

func rob(nums []int) int {
	result := 0
	if len(nums) == 1 {
		result = nums[0]
	} else if len(nums) == 2 {
		result = maxInt(nums[0], nums[1])
	} else if len(nums) == 3 {
		result = maxInt(nums[0], nums[1], nums[2])
	} else if len(nums) > 0 {
		savenums := []int{}
		savenums = append(savenums, nums...)
		savenums = append(savenums, nums...)
		save := make([]int, len(nums))
		for i := 0; i < len(save); i++ {
			tmpsave := make([]int, len(savenums))
			// 初始化最开始的数据
			tmpsave[i] = savenums[i]
			tmpsave[i+len(nums)] = savenums[i]
			tmpsave[i+1] = maxInt(savenums[i], savenums[i+1])
			tmpsave[i+2] = maxInt(savenums[i]+savenums[i+2], savenums[i+1])
			for j := i + 3; j < i+len(nums); j++ {
				// 处理最后一个数据的情况
				if tmpsave[j+1] == 0 {
					tmpsave[j] = maxInt(tmpsave[j-1], tmpsave[j-2]+savenums[j], tmpsave[j-3]+savenums[j])
				} else {
					tmpsave[j] = maxInt(tmpsave[j-1], tmpsave[j-2], tmpsave[j-3])
				}
			}
			//fmt.Println("i = ", i, "tmp save = ", tmpsave)
			if result < tmpsave[i+len(nums)-1] {
				result = tmpsave[i+len(nums)-1]
			}
		}
	}
	return result
}

func main() {

	nums := []int{1, 2, 3, 1}
	fmt.Println("result = ", rob(nums))

}

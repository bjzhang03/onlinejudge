package main

import (
	"fmt"
	"math"
	"sort"
)

// 堆调整
func headAdjust(nums *[]int, position int) {
	if position < len(*nums) {
		//fmt.Println("position = ", position, "nums = ", *nums)
		left := position * 2
		right := position*2 + 1

		if right < len(*nums) && (*nums)[position] < (*nums)[right] {
			tmp := (*nums)[position]
			(*nums)[position] = (*nums)[right]
			(*nums)[right] = tmp
		}

		if left < len(*nums) && (*nums)[position] < (*nums)[left] {
			tmp := (*nums)[position]
			(*nums)[position] = (*nums)[left]
			(*nums)[left] = tmp
		}
		// 对堆数据进行递归调整
		headAdjust(nums, left)
		headAdjust(nums, right)
	}
}

// 获取值最小的元素和他们的下标
func minNum(nums []int) (int, int) {
	result := math.MaxInt32
	position := 0
	for i := 1; i < len(nums); i++ {
		if result > nums[i] {
			result = nums[i]
			position = i
		}
	}
	return result, position
}

// 将元素交换带第一个数据的位置
func swapNum(nums *[]int, position int) {
	tmp := (*nums)[1]
	(*nums)[1] = (*nums)[position]
	(*nums)[position] = tmp
}

func nthSuperUglyNumber(n int, primes []int) int {
	result := 0
	if n > 0 && len(primes) > 0 {
		save := []int{}
		sort.Ints(primes)
		lengthMul := primes[len(primes)-1]/primes[0] + 2
		candidate := []int{0, 1}
		canused := map[int]bool{1: true}
		for len(save) < n {
			// 添加进来一个元素
			nextadd, minpos := minNum(candidate)
			// 将数据添加到save里面去
			save = append(save, nextadd)
			// 将数据更换到第一个的位置
			swapNum(&candidate, minpos)
			//标记数据已经使用过了
			candidate[1] = math.MaxInt32 - 1

			for i := 0; i < len(primes); i++ {
				if _, ok := canused[primes[i]*nextadd]; !ok {
					if len(candidate) < len(primes)*lengthMul {
						candidate = append(candidate, primes[i]*nextadd)
						for i := len(primes) - 1; i >= 1; i-- {
							headAdjust(&candidate, i)
						}
					} else if primes[i]*nextadd < candidate[1] {
						//始终保存最小的k个数据
						candidate[1] = primes[i] * nextadd
						delete(canused, candidate[1])
					}
					headAdjust(&candidate, 1)
					canused[primes[i]*nextadd] = true
				}
			}
			headAdjust(&candidate, 1)
			fmt.Println(save, candidate)
		}
		result = save[len(save)-1]

		//fmt.Println(save, candidate)
	}
	return result
}

func main() {
	primes := []int{37, 43, 59, 61, 67, 71, 79, 83, 89, 97, 101, 103, 113, 127, 131, 157, 163, 167, 173, 179, 191, 193, 197, 199, 211, 229, 233, 239, 251, 257}
	fmt.Println(nthSuperUglyNumber(800, primes))
}

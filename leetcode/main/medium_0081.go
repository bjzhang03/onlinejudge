package main

import "fmt"

func binarySearch(nums []int, target int, ins int, ine int) bool {
	result := false
	for ins <= ine {
		inmid := (ins + ine) / 2
		if nums[inmid] == target {
			result = true
			break
		} else if nums[inmid] > target {
			ine = inmid - 1
		} else if nums[inmid] < target {
			ins = inmid + 1
		}
	}
	return result
}

func search(nums []int, target int) bool {
	result := false
	if len(nums) > 0 {
		ins := 0             // index start
		ine := len(nums) - 1 // index end

		for ins <= ine {
			//fmt.Println("ins = ", ins, "ine = ", ine)
			inmid := (ins + ine) / 2 // index mid
			if nums[inmid] == target {
				result = true
				break
			} else if nums[inmid] > nums[ins] {
				// fmt.Println("front")
				// 前半部分是有序的
				if nums[inmid] <= target {
					// fmt.Println("front 1")
					ins = inmid + 1
				} else if nums[ins] <= target {
					// fmt.Println("front 2")
					result = binarySearch(nums, target, ins, inmid-1)
					break
				} else if nums[ins] > target {
					//fmt.Println("front 3")
					ins = inmid + 1
				}
			} else if nums[inmid] < nums[ins] {
				//fmt.Println("back")
				// 后半部分是有序的
				if nums[inmid] >= target {
					ine = inmid - 1
				} else if nums[ine] >= target {
					result = binarySearch(nums, target, inmid+1, ine)
					break
				} else if nums[ine] < target {
					ine = inmid - 1
				}
			} else if nums[inmid] == nums[ins] {
				// 特殊情况，无法进行有序无序的判断
				for i := ins; i <= ine; i++ {
					if nums[i] == target {
						result = true
						break
					}
				}
				break
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 1, 1, 1}
	fmt.Println(search(nums, 3))

}

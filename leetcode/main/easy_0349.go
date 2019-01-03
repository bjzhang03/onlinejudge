package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	if len(nums1) > 0 && len(nums2) > 0 {
		savenums1 := make(map[int]bool)
		savenums2 := make(map[int]bool)
		for i := 0; i < len(nums1); i++ {
			savenums1[nums1[i]] = true
		}
		for i := 0; i < len(nums2); i++ {
			if _, ok := savenums1[nums2[i]]; ok {
				// 去掉重复出现的数据
				if _, ok2 := savenums2[nums2[i]]; !ok2 {
					result = append(result, nums2[i])
					savenums2[nums2[i]] = true
				}
			}
		}
	}
	return result
}

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))

}

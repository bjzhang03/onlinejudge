package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > 0 && len(nums2) > 0 {
		// 建立map来进行计数操作
		savenums1 := make(map[int]int)
		savenums2 := make(map[int]int)
		for i := 0; i < len(nums1); i++ {
			if _, ok := savenums1[nums1[i]]; ok {
				savenums1[nums1[i]]++
			} else {
				savenums1[nums1[i]] = 1
			}
		}
		for i := 0; i < len(nums2); i++ {
			if _, ok := savenums2[nums2[i]]; ok {
				savenums2[nums2[i]]++
			} else {
				savenums2[nums2[i]] = 1
			}
		}
		result := []int{}
		for k, v := range savenums1 {
			if _, ok := savenums2[k]; ok {
				count := v
				if count > savenums2[k] {
					count = savenums2[k]
				}

				for i := 0; i < count; i++ {
					result = append(result, k)
				}
			}
		}
		return result
	}
	return []int{}
}

func main() {
	nums1 := []int{61, 24, 20, 58, 95, 53, 17, 32, 45, 85, 70, 20, 83, 62, 35, 89, 5, 95, 12, 86, 58, 77, 30, 64, 46, 13, 5, 92, 67, 40, 20, 38, 31, 18, 89, 85, 7, 30, 67, 34, 62, 35, 47, 98, 3, 41, 53, 26, 66, 40, 54, 44, 57, 46, 70, 60, 4, 63, 82, 42, 65, 59, 17, 98, 29, 72, 1, 96, 82, 66, 98, 6, 92, 31, 43, 81, 88, 60, 10, 55, 66, 82, 0, 79, 11, 81}
	nums2 := []int{5, 25, 4, 39, 57, 49, 93, 79, 7, 8, 49, 89, 2, 7, 73, 88, 45, 15, 34, 92, 84, 38, 85, 34, 16, 6, 99, 0, 2, 36, 68, 52, 73, 50, 77, 44, 61, 48}
	fmt.Println(intersect(nums1, nums2))

}

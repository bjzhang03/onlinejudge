package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := []int{}
	start1 := 0
	start2 := 0
	for ; start1 < m && start2 < n; {
		if nums1[start1] <= nums2[start2] {
			tmp = append(tmp, nums1[start1])
			start1++
		} else {
			tmp = append(tmp, nums2[start2])
			start2++
		}
	}

	for ; start1 < m; {
		tmp = append(tmp, nums1[start1])
		start1++
	}

	for ; start2 < n; {
		tmp = append(tmp, nums2[start2])
		start2++
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = tmp[i]
	}
}

func main() {

	nums1 := []int{1, 0}
	nums2 := []int{2}
	merge(nums1, 1, nums2, 1)

}

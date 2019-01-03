package main

type NumArray struct {
	save []int
}

func Constructor(nums []int) NumArray {
	result := &NumArray{[]int{}}
	for i := 0; i < len(nums); i++ {
		result.save = append(result.save, nums[i])
	}
	return *result
}

func (this *NumArray) SumRange(i int, j int) int {
	result := 0
	if i >= 0 && j < len(this.save) {
		for k := i; k <= j; k++ {
			result = result + this.save[k]
		}
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

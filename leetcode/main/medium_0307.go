package main

type NumArray struct {
	save []int
	sum  []int
}

func Constructor(nums []int) NumArray {
	obj := &NumArray{make([]int, len(nums)), make([]int, len(nums))}
	for i := 0; i < len(nums); i++ {
		obj.save[i] = nums[i]
		obj.sum[i] = nums[i]
		if i-1 >= 0 {
			obj.sum[i] = obj.sum[i] + obj.sum[i-1]
		}
	}
	return *obj
}

func (this *NumArray) Update(i int, val int) {
	if i < len(this.save) {
		this.save[i] = val
		// 更新和数据
		for j := i; j < len(this.save); j++ {
			this.sum[j] = this.save[j]
			if j > 0 {
				this.sum[j] = this.sum[j] + this.sum[j-1]
			}
		}
	}

}

func (this *NumArray) SumRange(i int, j int) int {
	result := 0
	if j < len(this.sum) {
		result = this.sum[j]
		if i-1 >= 0 {
			result = result - this.sum[i-1]
		}
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

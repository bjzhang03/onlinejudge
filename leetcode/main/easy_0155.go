package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	save []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	save := []int{}
	minStack := MinStack{save, math.MaxInt32}
	return minStack
}

func (this *MinStack) Push(x int) {
	this.save = append(this.save, x)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.save) > 0 {
		this.save = this.save[:len(this.save)-1]
		this.min = math.MaxInt32
		for i := 0; i < len(this.save); i++ {
			if this.min > this.save[i] {
				this.min = this.save[i]
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.save) > 0 {
		return this.save[len(this.save)-1]
	} else {
		return 0
	}
}

func (this *MinStack) GetMin() int {
	return this.min

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor();
	obj.Push(2);
	obj.Push(-1);
	obj.Pop();
	param_3 := obj.Top();
	param_4 := obj.GetMin();
	fmt.Println(param_3)
	fmt.Println(param_4)

}

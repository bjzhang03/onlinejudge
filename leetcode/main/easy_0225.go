package main

type MyStack struct {
	save []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	obj := &MyStack{[]int{}}
	return *obj
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.save = append(this.save, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	result := 0
	if len(this.save) > 0 {
		result = this.save[len(this.save)-1]
		this.save = this.save[:len(this.save)-1]
	}
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	result := 0
	if len(this.save) > 0 {
		result = this.save[len(this.save)-1]
	}
	return result
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	result := true
	if len(this.save) > 0 {
		result = false
	}
	return result

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}

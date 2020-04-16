package main

import (
	"fmt"
)

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())

	/**
	 * Your MinStack object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Push(x);
	 * obj.Pop();
	 * param_3 := obj.Top();
	 * param_4 := obj.GetMin();
	 */
}

type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	if curMin := this.GetMin(); x < curMin || len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, curMin)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
	if len(this.min) != 0 {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

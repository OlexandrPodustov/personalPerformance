package main

// Your MyQueue object will be instantiated and called as such:
// func main() {

// 	obj := Constructor()
// 	obj.Push(11)
// 	param_2 := obj.Pop()
// 	param_3 := obj.Peek()
// 	param_4 := obj.Empty()

// 	fmt.Println(param_2)
// 	fmt.Println(param_3)
// 	fmt.Println(param_4)
// }

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	res := this.stack[0]
	this.stack = this.stack[1:]

	return res
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

// LeetCode 232. Implement Queue using Stacks
package main

import "fmt"

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{stack: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	poped := this.stack[0]
	this.stack = this.stack[1:]
	return poped
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}

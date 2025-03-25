// LeetCode 232. Implement Queue using Stacks
package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	this.move()
	poped := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return poped
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func (this *MyQueue) move() {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			x := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, x)
		}
	}
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}

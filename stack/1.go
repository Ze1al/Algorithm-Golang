package stack

// 栈
// https://leetcode.cn/problems/implement-queue-using-stacks/

//type MyQueue struct {
//	inStack  []int
//	outStack []int
//}
//
//func Constructor() MyQueue {
//	return MyQueue{
//		inStack:  []int{},
//		outStack: []int{},
//	}
//}
//
//func (this *MyQueue) Push(x int) {
//	this.inStack = append(this.inStack, x)
//}
//
//func (this *MyQueue) Pop() int {
//	if len(this.outStack) == 0 {
//		for i := len(this.inStack) - 1; i >= 0; i-- {
//			this.outStack = append(this.outStack, this.inStack[i])
//		}
//		this.inStack = []int{}
//	}
//	res := this.outStack[len(this.outStack)-1]
//	this.outStack = this.outStack[:len(this.outStack)-1]
//	return res
//}
//
//func (this *MyQueue) Peek() int {
//	if len(this.outStack) == 0 {
//		for i := len(this.inStack) - 1; i >= 0; i-- {
//			this.outStack = append(this.outStack, this.inStack[i])
//		}
//		this.inStack = []int{}
//	}
//	return this.outStack[len(this.outStack)-1]
//}
//
//func (this *MyQueue) Empty() bool {
//	return len(this.outStack) == 0 && len(this.inStack) == 0
//}

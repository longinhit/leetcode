/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (38.09%)
 * Total Accepted:    121.1K
 * Total Submissions: 317.4K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 *
 *
 * Example:
 *
 *
 * MyStack stack = new MyStack();
 *
 * stack.push(1);
 * stack.push(2);
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 *
 * Notes:
 *
 *
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 *
 *
 */
// package main

// import "fmt"

// func main() {
// 	fmt.Printf("ddd")
// }

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if len(this.queue1) == 0 {
		this.queue1 = append(this.queue1, x)
		for _, temp := range this.queue2 {
			this.queue1 = append(this.queue1, temp)
		}
		this.queue2 = []int{}
	} else {
		this.queue2 = append(this.queue2, x)
		for _, temp := range this.queue1 {
			this.queue2 = append(this.queue2, temp)
		}
		this.queue1 = []int{}
	}

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	temp := 0
	if len(this.queue1) > 0 {
		temp = this.queue1[0]
		this.queue1 = this.queue1[1:]
	} else {
		temp = this.queue2[0]
		this.queue2 = this.queue2[1:]
	}
	return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queue1) > 0 {
		return this.queue1[0]
	} else {
		return this.queue2[0]
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

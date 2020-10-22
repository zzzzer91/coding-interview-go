// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

package main

func main() {

}

type MinStack struct {
	stack1 []int
	stack2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack1) == 0 || x <= this.stack1[len(this.stack1)-1] {
		this.stack1 = append(this.stack1, x)
	}
	this.stack2 = append(this.stack2, x)
}

func (this *MinStack) Pop() {
	if len(this.stack2) == 0 {
		return
	}
	n := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	if n == this.stack1[len(this.stack1)-1] {
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack2) == 0 {
		return -1
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack1) == 0 {
		return -1
	}
	return this.stack1[len(this.stack1)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

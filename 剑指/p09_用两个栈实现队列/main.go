// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

package main

func main() {

}

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
			return -1
		}
		stack := this.stack1
		this.stack1 = nil
		l, r := 0, len(stack)-1
		for l < r {
			stack[l], stack[r] = stack[r], stack[l]
			l++
			r--
		}
		this.stack2 = stack
	}
	tmp := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return tmp
}

// https://leetcode-cn.com/problems/min-stack/

package main

func main() {

}

type MinStack struct {
	mins []int
	vals []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.vals = append(this.vals, x)
	if len(this.mins) == 0 || x <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}

	if len(this.mins) != 0 {
		if this.vals[len(this.vals)-1] == this.mins[len(this.mins)-1] {
			this.mins = this.mins[:len(this.mins)-1]
		}
	}

	this.vals = this.vals[:len(this.vals)-1]
}

func (this *MinStack) Top() int {
	if len(this.vals) == 0 {
		return 0
	}

	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}

	return this.mins[len(this.mins)-1]
}

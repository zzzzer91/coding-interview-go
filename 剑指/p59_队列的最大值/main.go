// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

package main

import "container/list"

func main() {

}

type MaxQueue struct {
	q    *list.List
	maxQ *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:    list.New(),
		maxQ: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.maxQ.Len() == 0 {
		return -1
	}
	return this.maxQ.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	for this.maxQ.Len() > 0 && value > this.maxQ.Back().Value.(int) {
		this.maxQ.Remove(this.maxQ.Back())
	}
	this.q.PushBack(value)
	this.maxQ.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.q.Len() == 0 {
		return -1
	}
	v := this.q.Remove(this.q.Front()).(int)
	if v == this.maxQ.Front().Value.(int) { // 在这之前，根据队列性质，队首弹出元素都比 v 小
		this.maxQ.Remove(this.maxQ.Front())
	}
	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

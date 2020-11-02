package main

type ArrayQueue struct {
	data   []interface{}
	cap    int
	length int
	front  int // 队空：`front==rear`，队满：`front==(rear+1)%cap`，会损失一格
	rear   int
}

func New(cap int) ArrayQueue {
	return ArrayQueue{
		data:   make([]interface{}, cap),
		cap:    cap,
		length: 0,
		front:  0,
		rear:   0,
	}
}

func (q *ArrayQueue) Cap() int {
	return q.cap
}

func (q *ArrayQueue) Len() int {
	return q.length
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *ArrayQueue) IsFull() bool {
	return q.front == (q.rear+1)%q.cap
}

func (q *ArrayQueue) Push(e interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % q.cap
	return true
}

func (q *ArrayQueue) Peak() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	e := q.data[q.front]
	q.front = (q.front + 1) % q.cap
	return e, true
}

func main() {

}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushFront("2")
	queue.PushFront("1")
	queue.PushBack("3")
	queue.PushBack("4")
	// 相当于 dequeue
	queue.Remove(queue.Front())
	// 相当于 pop
	queue.Remove(queue.Back())
	fmt.Println(queue.Len())
	fmt.Println("*************************")
	fmt.Println(queue.Front().Value.(string))
	fmt.Println(queue.Back().Value.(string))
}

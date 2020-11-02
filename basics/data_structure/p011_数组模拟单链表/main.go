package main

import "fmt"

const N = 1e5 + 10

var e, ne [N]int  // e 相当于链表的节点的 Val，ne 相当于 链表节点的 Next
var head, idx int // idx 表示当前用到了哪个节点，head 存储链表头

func main() {
	initList()
	insertToHead(1)
	insertToHead(2)
	insertToHead(3)
	removeHead()
	insertToHead(4)
	traversal()
}

func initList() {
	head = -1 // 相当于 nil
	idx = 0
}

// 插在链表头
func insertToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// 删除头部
func removeHead() {
	if head != -1 {
		head = ne[head]
	}
}

func traversal() {
	for i := head; i != -1; i = ne[i] {
		fmt.Println(e[i])
	}
}

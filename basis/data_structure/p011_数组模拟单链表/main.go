package main

import "fmt"

const N = 1e5 + 10

var tail, idx int // idx 表示当前用到了哪个节点，tail 存储链表尾
var e, ne [N]int  // e 相当于链表的节点的 Val，ne 相当于 链表节点的 Next

func main() {
	initList()
	insert(1)
	insert(2)
	insert(3)
	remove()
	traversal()
}

func initList() {
	tail = -1 // 相当于 nil
	idx = 0
}

// 插在链表尾
func insert(a int) {
	e[idx] = a
	ne[idx] = tail
	tail = idx
	idx++
}

// 移除链表尾指向的节点
func remove() {
	if tail != -1 {
		tail = ne[tail]
	}
}

func traversal() {
	for i := 0; i <= tail; i++ {
		fmt.Println(e[i])
	}
}

// https://leetcode-cn.com/problems/lru-cache/

package main

import "container/list"

func main() {

}

type LRUCache struct {
	capacity int
	length   int
	cache    map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		cache:    make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.list.MoveToBack(v)
		return v.Value.([2]int)[1]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.Value = [2]int{key, value}
		this.list.MoveToBack(v)
	} else {
		if this.length >= this.capacity {
			v := this.list.Remove(this.list.Front()).([2]int)
			delete(this.cache, v[0])
		}
		this.cache[key] = this.list.PushBack([2]int{key, value})
		this.length++
	}
}

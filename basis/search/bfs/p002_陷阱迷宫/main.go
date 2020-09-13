// n*n 迷宫，出口在右下角，0 正常路，消耗 1s，# 代表陷阱，消耗 k s，1 代表墙，不能通行
// 求最短耗时
// 思路：优先级队列 bfs
/*
3 2
0#0
0#1
000
*/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 110

type Point struct {
	X    int
	Y    int
	Time int // 记录到当前点的最短时间
}

type MinHeap []Point

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Time < h[j].Time
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

var n, k int
var g [N][N]byte
var dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

func main() {
	fmt.Scanf("%d %d", &n, &k)
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		r.Read(g[i][:n+1]) // +1 是换行符
	}
	fmt.Println(g[:n])
	var h MinHeap
	heap.Push(&h, Point{0, 0, 0})
	g[0][0] = '1'     // 标记为访问过
	for h.Len() > 0 { // 优先级队列
		p := h[0]
		heap.Pop(&h)
		if p.X == n-1 && p.Y == n-1 { // 找到出口，如果出口不在右下角，可以替换为其他坐标
			fmt.Println(p.Time)
			return
		}
		for i := 0; i < 4; i++ {
			x, y := p.X+dx[i], p.Y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < n && g[x][y] != '1' {
				time := p.Time
				if g[x][y] == '#' {
					time += k
				} else if g[x][y] == '0' {
					time += 1
				}
				g[x][y] = '1' // 代表该位置已经被访问过了，下次不再访问
				heap.Push(&h, Point{x, y, time})
			}
		}
	}
}

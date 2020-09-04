// https://www.acwing.com/problem/content/846/
// bfs 能搜到最短路，dfs不能
// bfs 有固定模板：
// queue <- 初始值
// while queue不空
//     t <- 队头
//     拓展队头

package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	X int
	Y int
}

const N = 110

var n, m int
var g [N][N]int
var st [N][N]int // 代表有没有被搜过了，只有没被搜过的才是最短路
var dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&g[i][j])
			st[i][j] = -1
		}
	}
	st[0][0] = 0
	fmt.Println(bfs())
}

func bfs() int {
	q := list.New()
	q.PushBack(Point{0, 0})
	for q.Len() > 0 {
		p := q.Remove(q.Front()).(Point)
		for i := 0; i < 4; i++ {
			x, y := p.X+dx[i], p.Y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && st[x][y] == -1 {
				st[x][y] = st[p.X][p.Y] + 1
				q.PushBack(Point{x, y})
			}
		}
	}
	return st[n-1][m-1]
}

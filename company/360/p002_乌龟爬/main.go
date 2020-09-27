package main

import "fmt"

const N = 10010

// n 个点的 m 条带权无向边，求能所有能从 s 点到 t 点的路径中，某一条边的最小权值
var n, m, s, t int
var g [N][N]int
var dist [N]int // 记录每一个点距离第一个点的距离
var st [N]bool  // 记录该点的最短距离是否已经确定

func main() {
	fmt.Scanf("%d %d %d %d", &n, &m, &s, &t)
	for i := 1; i <= n; i++ {
		dist[i] = 1e9
	}
	dist[1] = 0
	for m > 0 {
		var u, v int
		fmt.Scanf("%d %d", &u, &v)
		fmt.Scanf("%d", &g[u][v])
		m--
	}
	fmt.Println(f())
}

func f() int {
	for i := 0; i < n; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		st[t] = true
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
	}
	return dist[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

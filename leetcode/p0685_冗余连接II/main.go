// https://leetcode-cn.com/problems/redundant-connection-ii/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
}

func findRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	numNodes := len(edges)
	uf := NewUnionFind(numNodes + 1)
	parent := make([]int, numNodes+1) // parent[i] 表示 i 的父节点
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.Find(from) == uf.Find(to) { // from 和 to 已连接
				cycleEdge = edge
			} else {
				uf.Union(from, to)
			}
		}
	}

	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func (uf UnionFind) Find(x int) int {
	if uf[x] != x {
		uf[x] = uf.Find(uf[x])
	}
	return uf[x]
}

func (uf UnionFind) Union(from, to int) {
	uf[uf.Find(from)] = uf.Find(to)
}

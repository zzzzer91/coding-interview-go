package main

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	uf := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i // 令根节点为集合编号
	}
	return uf
}

// 返回 x 的祖宗节点（即集合编号）
func (uf UnionFind) Find(x int) int {
	if uf[x] != x {
		uf[x] = uf.Find(uf[x]) // 路径压缩
	}
	return uf[x]
}

func (uf UnionFind) Union(from, to int) {
	uf[uf.Find(from)] = uf.Find(to)
}

func main() {

}

package main

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
		uf[x] = uf.Find(uf[x]) // 路径压缩
	}
	return uf[x]
}

func (uf UnionFind) Union(from, to int) {
	uf[uf.Find(from)] = uf.Find(to)
}

func main() {

}

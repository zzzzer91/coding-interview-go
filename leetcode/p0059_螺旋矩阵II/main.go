// https://leetcode-cn.com/problems/spiral-matrix-ii/

package main

func main() {

}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	count := 1
	ll, rr, tt, bb := 0, n-1, 0, n-1
	for ll <= rr && tt <= bb {
		for i := ll; i <= rr; i++ {
			res[tt][i] = count
			count++
		}
		for i := tt + 1; i <= bb; i++ {
			res[i][rr] = count
			count++
		}
		for i := rr - 1; i >= ll && tt < bb; i-- {
			res[bb][i] = count
			count++
		}
		for i := bb - 1; i > tt && ll < rr; i-- {
			res[i][ll] = count
			count++
		}
		ll++
		rr--
		tt++
		bb--
	}
	return res
}

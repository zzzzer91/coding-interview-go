// https://leetcode-cn.com/problems/spiral-matrix/

package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)
	ll, rr, tt, bb := 0, n-1, 0, m-1
	for ll <= rr && tt <= bb {
		for i := ll; i <= rr; i++ {
			res = append(res, matrix[tt][i])
		}
		for i := tt + 1; i <= bb; i++ {
			res = append(res, matrix[i][rr])
		}
		// && tt < bb 为了应对行数是奇数
		for i := rr - 1; i >= ll && tt < bb; i-- {
			res = append(res, matrix[bb][i])
		}
		// && ll < rr 为了应对列数是奇数
		for i := bb - 1; i > tt && ll < rr; i-- {
			res = append(res, matrix[i][ll])
		}
		ll++
		rr--
		tt++
		bb--
	}
	return res
}

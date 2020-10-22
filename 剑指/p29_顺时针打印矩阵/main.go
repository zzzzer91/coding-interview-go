// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	colLen, rowLen := len(matrix), len(matrix[0])
	res := make([]int, 0, colLen*rowLen)
	tt, bb, ll, rr := 0, colLen-1, 0, rowLen-1
	for tt <= bb && ll <= rr {
		for i := ll; i <= rr; i++ {
			res = append(res, matrix[tt][i])
		}
		for i := tt + 1; i <= bb; i++ {
			res = append(res, matrix[i][rr])
		}
		for i := rr - 1; i >= ll && tt < bb; i-- {
			res = append(res, matrix[bb][i])
		}
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

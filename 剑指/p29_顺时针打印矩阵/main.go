// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)
	tt, bb, ll, rr := 0, m-1, 0, n-1
	for tt <= bb && ll <= rr {
		for i := ll; i <= rr; i++ {
			res = append(res, matrix[tt][i])
		}
		for i := tt + 1; i <= bb; i++ {
			res = append(res, matrix[i][rr])
		}
		// tt < bb 这个条件要加上，否则迭代出错的例子：
		// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
		for i := rr - 1; i >= ll && tt < bb; i-- {
			res = append(res, matrix[bb][i])
		}
		// ll < rr 这个条件要加上，否则迭代出错的例子：
		// [[7],[9],[6]]
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

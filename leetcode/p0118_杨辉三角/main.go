// https://leetcode-cn.com/problems/pascals-triangle/

package main

func main() {

}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, 0, numRows)
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		preRow := res[i-1]
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = preRow[j-1] + preRow[j]
		}
		res = append(res, row)
	}
	return res
}

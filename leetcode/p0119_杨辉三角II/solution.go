package p0119_杨辉三角II

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			ret[j] += ret[j-1]
		}
	}
	return ret
}

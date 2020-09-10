package main

func main() {

}

func generate(numRows int) [][]int {
	var ret [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < i; j++ {
			preRow := ret[i-1]
			row[j] = preRow[j-1] + preRow[j]
		}
		ret = append(ret, row)
	}
	return ret
}

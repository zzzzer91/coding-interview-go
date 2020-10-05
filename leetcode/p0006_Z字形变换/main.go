// https://leetcode-cn.com/problems/zigzag-conversion/

package main

func main() {

}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	temp := make([][]byte, numRows)
	pos := 0
	flag := -1
	for i := range s {
		temp[pos] = append(temp[pos], s[i])
		if pos == 0 || pos == numRows-1 {
			flag = -flag
		}
		pos += flag
	}
	var res string
	for _, v := range temp {
		res += string(v)
	}
	return res
}

// https://leetcode-cn.com/problems/integer-to-roman/

package main

func main() {

}

func intToRoman(num int) string {
	intSlice := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	for num > 0 {
		for i, n := range intSlice {
			if num >= n {
				res += roman[i]
				num -= n
				break
			}
		}
	}
	return res
}

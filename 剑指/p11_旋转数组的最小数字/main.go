// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

package main

func main() {

}

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		m := l + (r-l)>>1
		if numbers[m] > numbers[r] {
			l = m + 1
		} else if numbers[m] < numbers[r] {
			r = m
		} else {
			r--
		}
	}
	return numbers[l]
}

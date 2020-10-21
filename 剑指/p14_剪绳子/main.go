// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/

package main

func main() {

}

// 数学推导，
// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/solution/mian-shi-ti-14-i-jian-sheng-zi-tan-xin-si-xiang-by/
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return pow(3, a)
	}
	if b == 1 {
		return pow(3, a-1) * 4
	}
	// 只可能 b == 2
	return pow(3, a) * 2
}

func pow(x, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= x
	}
	return res
}

// https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/

package main

func main() {

}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res = (res * 3) % (1e9 + 7)
		n -= 3
	}
	return (res * n) % (1e9 + 7)
}

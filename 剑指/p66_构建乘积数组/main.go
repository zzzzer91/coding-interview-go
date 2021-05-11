// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/

package main

func main() {

}

func constructArr(a []int) []int {
	res := make([]int, len(a))
	for i := range res {
		res[i] = 1
	}
	lv, rv := 1, 1
	for i := range a {
		res[i] *= lv
		lv *= a[i]
		res[len(a)-i-1] *= rv
		rv *= a[len(a)-i-1]
	}
	return res
}

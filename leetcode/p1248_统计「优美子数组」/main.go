// https://leetcode-cn.com/submissions/detail/174215144/

package main

func main() {

}

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	m := make([]int, len(nums)+1)
	m[0] = 1
	oddCount := 0
	for _, n := range nums {
		if n&1 != 0 {
			oddCount++
		}
		if oddCount >= k {
			res += m[oddCount-k]
		}
		m[oddCount]++
	}
	return res
}

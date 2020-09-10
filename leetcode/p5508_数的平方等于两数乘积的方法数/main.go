// 观察数据范围，基本可以得出结论，算法复杂度为 log(n^2)
// 所以要用二重循环
// 利用哈希表将三重循环优化为二重循环

package main

func main() {

}

func numTriplets(nums1 []int, nums2 []int) int {
	return f(nums1, nums2) + f(nums2, nums1)
}

func f(a1, a2 []int) int {
	res := 0
	m := make(map[int]int)
	for i := range a1 {
		m[a1[i]*a1[i]]++
	}
	for j := 0; j < len(a2); j++ {
		for k := j + 1; k < len(a2); k++ {
			res += m[a2[j]*a2[k]]
		}
	}
	return res
}

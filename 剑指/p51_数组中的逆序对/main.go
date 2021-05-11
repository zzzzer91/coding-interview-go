// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

package main

func main() {

}

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	m := (len(nums) - 1) >> 1
	// 总逆序对数量为：左边区间的逆序对+右边区间的逆序对+横跨两个区间的逆序对
	res := reversePairs(nums[:m+1]) + reversePairs(nums[m+1:])
	if nums[m] <= nums[m+1] {
		return res
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	i, j, k := 0, m+1, 0
	for i <= m && j < len(nums) {
		if temp[i] <= temp[j] { // 注意这里必须是小于等于
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
			res += m - i + 1
		}
		k++
	}
	if i <= m {
		copy(nums[k:], temp[i:m+1])
	} else {
		copy(nums[k:], temp[j:])
	}
	return res
}

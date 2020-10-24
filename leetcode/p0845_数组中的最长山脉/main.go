// https://leetcode-cn.com/problems/longest-mountain-in-array/

package main

func main() {

}

// 动态规划，枚举山顶，时O(n)，空O(n)
func longestMountain(a []int) (ans int) {
	n := len(a)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if a[i-1] < a[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if a[i+1] < a[i] {
			right[i] = right[i+1] + 1
		}
	}
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = l + r + 1
		}
	}
	return
}

// 双指针，枚举山脚，时O(n)，空O(1)
func longestMountain2(a []int) (ans int) {
	n := len(a)
	left := 0
	for left+2 < n {
		right := left + 1
		if a[left] < a[left+1] {
			for right+1 < n && a[right] < a[right+1] {
				right++
			}
			if right < n-1 && a[right] > a[right+1] {
				for right+1 < n && a[right] > a[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return
}

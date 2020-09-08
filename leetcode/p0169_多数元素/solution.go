package main

func main() {

}

// hash 表，时 O(n)，空 O(n)，虽然显示空间小于 100%，但非最佳
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	border := len(nums) / 2
	for k, v := range m {
		if v > border {
			return k
		}
	}
	return -1
}

// 多数投票算法(Boyer-Moore Algorithm)，时 O(n)，空 O(1)
// 注意：题目中已经表明必定存在一个数出现cissus大于 len(nums) / 2，
// 所以这里只需要扫描一次，如果没有这个条件，那么还需扫描第二次，来判断是不是众数
func majorityElement2(nums []int) int {
	var target, count int
	for _, num := range nums {
		if count == 0 {
			target = num
		}
		if target == num {
			count++
		} else {
			count--
		}
	}
	return target
}

package main

func main() {

}

func longestConsecutive(nums []int) int {
	ret := 0
	dict := make(map[int]bool, len(nums))
	for _, n := range nums {
		dict[n] = true
	}
	for _, n := range nums {
		// 如果前面还有连续的，则跳过当前的
		if !dict[n-1] {
			length := 1
			// 不停的往后找连续的
			for i := n + 1; dict[i]; i++ {
				length++
			}
			// 与之前长度比对
			ret = max(ret, length)
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

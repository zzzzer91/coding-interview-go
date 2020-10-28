// https://leetcode-cn.com/problems/unique-number-of-occurrences/

package main

func main() {

}

//根据要求，数字在-1000到1000之间，设置一个2000的数组,遍历统计次数
//然后由于arr在1000以内，设置另一个数组，记录次数，如果次数重复出现，则返回false
func uniqueOccurrences(arr []int) bool {
	counts := make([]int, 2000) // 数组模拟字典
	for _, v := range arr {
		i := v + 1000
		counts[i]++
	}
	freqs := make([]bool, 1000) // 数组模拟集合
	for _, v := range counts {
		if v == 0 {
			continue
		}
		if freqs[v] {
			return false
		}
		freqs[v] = true
	}
	return true
}

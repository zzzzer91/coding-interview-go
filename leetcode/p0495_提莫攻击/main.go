// https://leetcode-cn.com/problems/teemo-attacking/

package main

func main() {

}

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	res := duration
	pre := timeSeries[0] + duration
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] > pre {
			res += duration
			pre = timeSeries[i] + duration
			continue
		}
		sum := timeSeries[i] + duration
		if sum > pre {
			res += sum - pre
			pre = sum
		}
	}
	return res
}

func findPoisonedDuration2(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	res := duration
	for i := 1; i < len(timeSeries); i++ {
		res += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

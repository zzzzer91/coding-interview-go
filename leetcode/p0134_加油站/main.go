// https://leetcode-cn.com/problems/gas-station/

package main

import "math"

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	min := math.MaxInt64
	minIndex := 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < min {
			min = sum
			minIndex = i
		}
	}
	if sum < 0 {
		return -1
	}
	return (minIndex + 1) % len(gas) // 找一个 sum 最小的地方的下一个点作为起点
}

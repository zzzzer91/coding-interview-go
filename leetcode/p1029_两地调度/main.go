// https://leetcode-cn.com/problems/two-city-scheduling/

package main

import "sort"

func main() {

}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	n := len(costs) / 2
	res := 0
	for i := 0; i < n; i++ {
		res += costs[i][0] + costs[i+n][1]
	}
	return res
}

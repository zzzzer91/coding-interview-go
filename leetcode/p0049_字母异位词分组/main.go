// https://leetcode-cn.com/problems/group-anagrams/

package main

import "sort"

func main() {

}

// 排序，效率差
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		tmp := format(s)
		m[tmp] = append(m[tmp], s)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func format(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

// 前26个质数
var list = [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

// 建立单词到数字（质数）的映射
func groupAnagrams2(strs []string) [][]string {
	m := map[int]int{}
	index := 0
	res := make([][]string, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		product := helper(strs[i])
		if k, ok := m[product]; ok {
			res[k] = append(res[k], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			m[product] = index
			index++
		}
	}
	return res
}

func helper(a string) int {
	res := 1
	for i := range a {
		res *= list[a[i]-'a']
	}
	return res
}

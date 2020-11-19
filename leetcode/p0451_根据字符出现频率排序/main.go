// https://leetcode-cn.com/problems/sort-characters-by-frequency/

package main

import "sort"

func main() {

}

type Pair struct {
	C     byte
	Count int
}

// 排序法
func frequencySort(s string) string {
	m := [128]int{}
	for i := range s {
		if s[i] > 0 && s[i] < 128 {
			m[s[i]]++
		}
	}
	var arr []Pair
	for i := 0; i < 128; i++ {
		if m[i] != 0 {
			arr = append(arr, Pair{C: byte(i), Count: m[i]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		if a.Count == b.Count {
			return a.C < b.C
		}
		return a.Count > b.Count
	})
	var res []byte
	for _, v := range arr {
		for i := 0; i < v.Count; i++ {
			res = append(res, v.C)
		}
	}
	return string(res)
}

// 桶排序思想，更好
func frequencySort2(s string) string {
	m := [128]int{}
	for i := range s {
		if s[i] > 0 && s[i] < 128 {
			m[s[i]]++
		}
	}
	buckets := make([][]byte, len(s)+1) // 某个字符出现频率最多为 len(s)
	for i := 0; i < 128; i++ {
		if m[i] != 0 {
			buckets[m[i]] = append(buckets[m[i]], byte(i))
		}
	}
	var res []byte
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, c := range buckets[i] {
			for j := 0; j < i; j++ {
				res = append(res, c)
			}
		}
	}
	return string(res)
}

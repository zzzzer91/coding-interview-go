// https://leetcode-cn.com/problems/find-common-characters/

package main

import "math"

func main() {

}

func commonChars(a []string) (ans []string) {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range a {
		freq := [26]int{} // 相当于一个集合
		for _, b := range word {
			freq[b-'a']++
		}
		// 集合求交集
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}

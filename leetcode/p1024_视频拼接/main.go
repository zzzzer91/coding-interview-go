// https://leetcode-cn.com/problems/video-stitching/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(videoStitching([][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9))
}

func videoStitching(clips [][]int, t int) (ans int) {
	maxn := make([]int, t)
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < t && r > maxn[l] {
			maxn[l] = r
		}
	}
	last, pre := 0, 0
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return
}

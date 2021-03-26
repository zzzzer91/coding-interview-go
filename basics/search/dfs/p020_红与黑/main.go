package main

import (
	"fmt"
)

const N = 21

var w, h int
var a [N][]byte

func main() {
	var initX, initY int
	for {
		fmt.Scan(&w, &h)
		if w == 0 && h == 0 {
			break
		}
		for i := 0; i < h; i++ {
			fmt.Scan(&a[i]) // go 中读取单个字符很麻烦，只能一行读取了
			for j := 0; j < w; j++ {
				if a[i][j] == '@' {
					initX, initY = i, j
				}
			}
		}
		fmt.Println(f(initX, initY))
	}
}

func f(x, y int) int {
	if x < 0 || x >= h || y < 0 || y >= w {
		return 0
	}
	if a[x][y] == '#' {
		return 0
	}
	a[x][y] = '#'
	return 1 + f(x-1, y) + f(x+1, y) + f(x, y-1) + f(x, y+1)
}

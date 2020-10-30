package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dx, dy := [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}
	m, n := len(image), len(image[0])
	originalColor := image[sr][sc]
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if image[x][y] != originalColor {
			return
		}
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < m && b >= 0 && b < n {
				dfs(a, b)
			}
		}
	}
	if originalColor != newColor {
		dfs(sr, sc)
	}
	return image
}

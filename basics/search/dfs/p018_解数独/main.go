// https://leetcode-cn.com/problems/sudoku-solver/

package main

import "fmt"

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	for _, row := range board {
		fmt.Printf("%s\n", row)
	}
}

func solveSudoku(board [][]byte) {
	const rowLen, colLen = 9, 9
	var rowSt [rowLen][10]bool
	var colSt [colLen][10]bool
	var cellSt [rowLen / 3][colLen / 3][10]bool
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			c := board[i][j]
			if c != '.' {
				n := c - '0'
				rowSt[i][n] = true
				colSt[j][n] = true
				cellSt[i/3][j/3][n] = true
			}
		}
	}
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if y == colLen {
			y = 0
			x++
		}
		if x == rowLen {
			return true
		}
		if board[x][y] == '.' {
			for i := byte(1); i <= 9; i++ {
				if !rowSt[x][i] && !colSt[y][i] && !cellSt[x/3][y/3][i] {
					board[x][y] = i + '0'
					rowSt[x][i] = true
					colSt[y][i] = true
					cellSt[x/3][y/3][i] = true
					if dfs(x, y+1) {
						return true
					}
					board[x][y] = '.'
					rowSt[x][i] = false
					colSt[y][i] = false
					cellSt[x/3][y/3][i] = false
				}
			}
		} else {
			return dfs(x, y+1)
		}
		return false
	}
	dfs(0, 0)
}

package piscine

import (
	"ex08/test/ft"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafe(board [8]int) bool {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if board[i] == board[j] || abs(board[i]-board[j]) == j-i {
				return false
			}
		}
	}
	return true
}

func setBoard(board [8]int, i int) {
	if i == 8 {
		if isSafe(board) {
			for _, v := range board {
				ft.PrintRune(rune(v) + '1')
			}
			ft.PrintRune('\n')
		}
		return
	}
	for j := 0; j < 8; j++ {
		board[i] = j
		setBoard(board, i+1)
	}
}

func EightQueens() {
	var board [8]int
	setBoard(board, 0)
}

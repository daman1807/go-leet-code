package main

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		memory := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				_, ok := memory[board[i][j]]
				if ok {
					return false
				}
				memory[board[i][j]] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		memory := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if string(board[j][i]) != "." {
				_, ok := memory[board[j][i]]
				if ok {
					return false
				}
				memory[board[j][i]] = true
			}
		}
	}

	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		boxes[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			box := i/3*3 + j/3
			if string(board[i][j]) != "." {
				_, ok := boxes[box][board[i][j]]
				if ok {
					return false
				}
				boxes[box][board[i][j]] = true
			}
		}
	}

	return true
}

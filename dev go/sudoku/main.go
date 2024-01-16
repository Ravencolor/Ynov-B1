package main

import (
	"fmt"
	"os"
)

const N = 9
const Empty = '.'

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < N; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == Empty {
				for num := byte('1'); num <= byte('9'); num++ { 
					if isValid(board, i, j, num) {
						board[i][j] = num
						if solveSudoku(board) {
							return true
						}
						board[i][j] = Empty
					}
				}
				return false
			}
		}
	}
	return true 
}

func printBoard(board [][]byte) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 10 {
		fmt.Println("Error")
		return
	}
	board := make([][]byte, N)
	for i := 0; i < N; i++ {
		board[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			arg := os.Args[i+1]
			if len(arg) != N {
				fmt.Println("Error")
				return
			}
			board[i][j] = arg[j]
		}
	}
	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

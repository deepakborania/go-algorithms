package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if board == nil || word == "" {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, word string, r, c int) bool {
	if len(word) == 0 {
		return true
	}
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) || board[r][c] != word[0] {
		return false
	}
	tmp := board[r][c]
	board[r][c] = '#'
	remWord := word[1:]
	found := helper(board, remWord, r+1, c) ||
		helper(board, remWord, r-1, c) ||
		helper(board, remWord, r, c-1) ||
		helper(board, remWord, r, c+1)

	board[r][c] = tmp
	return found
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
}

package main

import "fmt"

// LeetCode: 51
// Place nqueens on n x n board in non-attacking position and return all distinct positions as follows:

// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
func solveNQueens(n int) [][]string {
	board := make([][]bool, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, n)
	}
	solutions := [][]string{}
	solveNQ(board, 0, n, &solutions)
	return solutions
}

// n = number of queens to be placed
// k = number of queens placed
func solveNQ(board [][]bool, k, n int, solutions *[][]string) {
	if k == n {
		//process solution
		sol := []string{}
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				if board[i][j] {
					str += "Q"
				} else {
					str += "."
				}
			}
			sol = append(sol, str)
		}
		*solutions = append(*solutions, sol)
		return
	}
	for i := 0; i < n; i++ {
		//Will board[k][n] work as solution?
		if isValid(board, k, i) {
			// Then use it
			board[k][i] = true
			solveNQ(board, k+1, n, solutions)
			// Then remove it
			board[k][i] = false
		}
	}
}

func isValid(board [][]bool, r, c int) bool {
	N := len(board)
	for i := 0; i < N; i++ {
		if board[r][i] || board[i][c] {
			return false
		}
	}
	//Lower right diagonal
	for i, j := r, c; i < N && j < N; i, j = i+1, j+1 {
		if board[i][j] {
			return false
		}
	}
	//Upper right diagonal
	for i, j := r, c; i >= 0 && j < N; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}
	//Upper left diagonal
	for i, j := r, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	//Lower left diagonal
	for i, j := r, c; i < N && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}
	return true
}

func main() {
	solutions := solveNQueens(4)
	for _, sol := range solutions {
		for _, c := range sol {
			fmt.Println(c)
		}
		fmt.Println("")
	}

}

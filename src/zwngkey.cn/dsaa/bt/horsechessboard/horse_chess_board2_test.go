/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 04:24:44
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 04:28:12
 */
package bt

import (
	"fmt"
	"testing"
)

const size = 8 // Size of the chessboard

var dx = []int{2, 1, -1, -2, -2, -1, 1, 2} // Possible x-coordinate offsets for knight moves
var dy = []int{1, 2, 2, 1, -1, -2, -2, -1} // Possible y-coordinate offsets for knight moves
var board [][]int                          // Chessboard to store the moves of the knight

func solve(x int, y int, move int) bool {
	if move == size*size {
		// All squares have been visited, so a solution has been found
		return true
	}

	// Try all possible knight moves from the current position
	for i := 0; i < 8; i++ {
		x1 := x + dx[i]
		y1 := y + dy[i]

		// Check if the move is valid
		if x1 >= 0 && x1 < size && y1 >= 0 && y1 < size && board[x1][y1] == 0 {
			board[x1][y1] = move
			if solve(x1, y1, move+1) {
				// The move led to a solution, so return true
				return true
			} else {
				// The move did not lead to a solution, so backtrack and try the next move
				board[x1][y1] = 0
			}
		}
	}
	// All possible moves from the current position have been tried, and none of them
	// led to a solution, so return false
	return false
}

func Test(t *testing.T) {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}

	// Start the knight at position (0, 0) on the chessboard
	board[0][0] = 1
	if solve(0, 0, 2) {
		// A solution was found, so print the moves of the knight on the chessboard
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				fmt.Printf("%2d ", board[i][j])
			}
			fmt.Println()
		}
	} else {
		// No solution was found
		fmt.Println("No solution was found.")
	}
}

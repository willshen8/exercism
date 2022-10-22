package stateoftictactoe

import (
	"fmt"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	statusBoard := strings.Split(strings.Join(board, ""), "")
	if DetermineInvalidGame(statusBoard) {
		return "", fmt.Errorf("invalid game")
	}

	if DetermineWin("X", statusBoard) || DetermineWin("O", statusBoard) {
		return Win, nil
	}
	if DetermineDrawGame(statusBoard) {
		return Draw, nil
	}
	return Ongoing, nil
}

// X starts, so number of X's should be equal to number of O's or 1 more
func DetermineInvalidGame(statusBoard []string) bool {
	var xCount, oCount int
	for _, v := range statusBoard {
		if v == "X" {
			xCount++
		}
		if v == "O" {
			oCount++
		}
	}
	if xCount-oCount > 1 || oCount-xCount >= 1 {
		return true
	}
	// determine if both wins
	if DetermineWin("X", statusBoard) && DetermineWin("O", statusBoard) {
		return true
	}
	return false
}

func DetermineDrawGame(statusBoard []string) bool {
	var xCount, oCount int
	for _, v := range statusBoard {
		if v == "X" {
			xCount++
		}
		if v == "O" {
			oCount++
		}
	}
	return xCount == 5 && oCount == 4
}

func DetermineWin(player string, board []string) bool {
	// Horizontal
	if board[0] == player && board[1] == player && board[2] == player {
		return true
	}
	if board[3] == player && board[4] == player && board[5] == player {
		return true
	}
	if board[6] == player && board[7] == player && board[8] == player {
		return true
	}

	// Vertical
	if board[0] == player && board[3] == player && board[6] == player {
		return true
	}
	if board[1] == player && board[4] == player && board[7] == player {
		return true
	}
	if board[2] == player && board[5] == player && board[8] == player {
		return true
	}

	// Diagonal
	if board[0] == player && board[4] == player && board[8] == player {
		return true
	}
	if board[2] == player && board[4] == player && board[6] == player {
		return true
	}

	return false
}

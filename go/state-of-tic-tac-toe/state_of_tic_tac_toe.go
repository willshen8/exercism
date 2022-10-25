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
	statusBoard := []rune(strings.Join(board, ""))
	hasXWon := isWinForPlayer('X', statusBoard)
	hasOWon := isWinForPlayer('O', statusBoard)

	if isInvalidGame(statusBoard) {
		return "", fmt.Errorf("invalid game")
	}

	if hasXWon && hasOWon {
		return "", fmt.Errorf("games has been played on after it was won")
	}

	if hasXWon || hasOWon {
		return Win, nil
	}
	if isDrawGame(statusBoard) {
		return Draw, nil
	}
	return Ongoing, nil
}

// X starts, so number of X's should be equal to number of O's or 1 more
func isInvalidGame(statusBoard []rune) bool {
	var xCount, oCount int
	for _, v := range statusBoard {
		if v == 'X' {
			xCount++
		}
		if v == 'O' {
			oCount++
		}
	}
	if xCount-oCount > 1 || oCount-xCount >= 1 {
		return true
	}
	// determine if both wins
	if isWinForPlayer('X', statusBoard) && isWinForPlayer('O', statusBoard) {
		return true
	}
	return false
}

func isDrawGame(statusBoard []rune) bool {
	var xCount, oCount int
	for _, v := range statusBoard {
		if v == 'X' {
			xCount++
		}
		if v == 'O' {
			oCount++
		}
	}
	return xCount == 5 && oCount == 4
}

func isWinForPlayer(player rune, board []rune) bool {
	return isWinHorizontally(player, board) || isWinVertically(player, board) || IsWinDiagonally(player, board)
}

func isWinHorizontally(player rune, board []rune) bool {
	if board[0] == player && board[1] == player && board[2] == player {
		return true
	}
	if board[3] == player && board[4] == player && board[5] == player {
		return true
	}
	if board[6] == player && board[7] == player && board[8] == player {
		return true
	}
	return false
}

func isWinVertically(player rune, board []rune) bool {
	if board[0] == player && board[3] == player && board[6] == player {
		return true
	}
	if board[1] == player && board[4] == player && board[7] == player {
		return true
	}
	if board[2] == player && board[5] == player && board[8] == player {
		return true
	}
	return false
}

func IsWinDiagonally(player rune, board []rune) bool {
	if board[0] == player && board[4] == player && board[8] == player {
		return true
	}
	if board[2] == player && board[4] == player && board[6] == player {
		return true
	}
	return false
}

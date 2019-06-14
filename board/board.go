package board

import (
	"errors"
	"fmt"
	"io"

	"github.com/gayanhewa/connect-four/players"
)

// Board represents the board.
type Board struct {
	grid  [6][7]players.Player
	moves []Move
}

// FirstMove returns true if this is the first move.
func (b *Board) FirstMove() bool {
	return len(b.moves) == 0
}

// Print the board to stdout.
func (b *Board) Print(w io.Writer) {
	fmt.Print("\n")
	fmt.Fprintln(w, "======================")
	for i := 5; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			p := b.grid[i][j]
			if p != nil {
				fmt.Fprintf(w, " %s ", string(p.Name()[0]))
			}
			if p == nil {
				fmt.Fprint(w, " ¤ ")
			}
		}
		fmt.Fprintln(w, "")
	}
	for j := 0; j < 7; j++ {
		fmt.Fprintf(w, " %d ", j)
	}
	fmt.Fprintln(w, "\n ======================")
}

//Move is an action by the player.
func (b *Board) Move(column int, player *players.Player) bool {
	for i := 0; i < 6; i++ {
		if b.grid[i][column] == nil {
			b.grid[i][column] = *player
			b.moves = append(b.moves, Move{
				column: column,
				player: player,
			})
			return true
		}
	}
	return false
}

// AvailableMoves prints available columns for a drop.
func (b *Board) AvailableMoves() []int {
	var moves []int
	for column := 0; column < 7; column++ {
		if b.grid[5][column] == nil {
			moves = append(moves, column)
		}
	}
	return moves
}

// IsValidMove checks if a move is valid.
func (b *Board) IsValidMove(column int) bool {
	return b.grid[5][column] == nil
}

// Winner returns a winning player or an error.
func (b *Board) Winner() (players.Player, error) {
	// Horizontal wins
	for row := 0; row < 6; row++ {
		for column := 0; column < 4; column++ {
			if b.grid[row][column] != nil && b.grid[row][column] == b.grid[row][column+1] && b.grid[row][column+1] == b.grid[row][column+2] && b.grid[row][column+2] == b.grid[row][column+3] {
				return b.grid[row][column], nil
			}
		}
	}
	// Vertical wins
	for column := 0; column < 7; column++ {
		for row := 0; row < 3; row++ {
			if b.grid[row][column] != nil && b.grid[row][column] == b.grid[row+1][column] && b.grid[row+1][column] == b.grid[row+2][column] && b.grid[row+2][column] == b.grid[row+3][column] {
				return b.grid[row][column], nil
			}
		}
	}

	// Diagnal wins from left to right
	for column := 0; column < 4; column++ {
		for row := 0; row < 3; row++ {
			if b.grid[row][column] != nil && b.grid[row][column] == b.grid[row+1][column+1] && b.grid[row+1][column+1] == b.grid[row+2][column+2] && b.grid[row+2][column+2] == b.grid[row+3][column+3] {
				return b.grid[row][column], nil
			}
		}
	}

	// Diagnal wins from right to left
	for column := 6; column > 2; column-- {
		for row := 0; row < 3; row++ {
			if b.grid[row][column] != nil && b.grid[row][column] == b.grid[row+1][column-1] && b.grid[row+1][column-1] == b.grid[row+2][column-2] && b.grid[row+2][column-2] == b.grid[row+3][column-3] {
				return b.grid[row][column], nil
			}
		}
	}
	return nil, errors.New("no winner yet")
}

// LastMove returns the last move played.
func (b *Board) LastMove() Move {
	return b.moves[len(b.moves)-1]
}

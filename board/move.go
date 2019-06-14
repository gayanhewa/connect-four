package board

import "github.com/gayanhewa/connect-four/players"

// Move store the move data.
type Move struct {
	column int
	player *players.Player
}

//Player returns the player responsible for the current move.
func (m *Move) Player() *players.Player {
	return m.player
}

// Column returns the column selected for this move.
func (m *Move) Column() int {
	return m.column
}

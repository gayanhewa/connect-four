package board

import (
	"testing"
)

func TestPlayerMove(t *testing.T) {
	player := humanPlayerHelper("Gayan")
	move := &Move{
		column: 5,
		player: &player,
	}
	if move.Column() != 5 {
		t.Fatal("Failed geting column from the move.")
	}
	if *move.Player() != player {
		t.Fatal("Failed geting player name from the move.")
	}
}

package players

import (
	"fmt"
	"strconv"

	"github.com/gayanhewa/connect-four/interactions"
)

// Human player.
type Human struct {
	name string
}

// NewHumanPlayer factory method.
func NewHumanPlayer(i interactions.Interactor) Player {
	playerName := i.Prompt("Enter your name to continue:")
	return &Human{
		playerName,
	}
}

// Name of the player.
func (h *Human) Name() string {
	return h.name
}

// Move plays the turn.
func (h *Human) Move(i interactions.Interactor, availableMoves []int) (int, error) {
	move := i.Prompt(fmt.Sprintf("%s enter your move [%v] :", h.Name(), availableMoves))
	parsedMove, err := strconv.Atoi(move)
	if err != nil {
		return -1, fmt.Errorf("invalid move to column %s", move)
	}
	if parsedMove < -1 || parsedMove > 6 {
		return -1, fmt.Errorf("invalid move to column %s", move)
	}
	if !h.isAvailableMove(parsedMove, availableMoves) {
		return -1, fmt.Errorf("invalid move to column %s", move)
	}
	return parsedMove, nil
}

func (h *Human) isAvailableMove(parsedMove int, availableMoves []int) bool {
	var exists bool
	for _, value := range availableMoves {
		if parsedMove == value {
			exists = true
		}
	}
	return exists
}

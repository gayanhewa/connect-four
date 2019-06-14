package players

import (
	"fmt"
	"math/rand"

	"github.com/gayanhewa/connect-four/interactions"
)

// ComputerPlayer player.
type ComputerPlayer struct {
	name string
}

// NewComputerPlayer factory method.
func NewComputerPlayer() Player {
	return &ComputerPlayer{
		name: fmt.Sprintf("%d-Computer", rand.Intn(9)),
	}
}

// Name of the player.
func (c *ComputerPlayer) Name() string {
	return c.name
}

// Move plays the turn.
func (c *ComputerPlayer) Move(i interactions.Interactor, availableMoves []int) (int, error) {
	return availableMoves[rand.Intn(len(availableMoves))], nil
}

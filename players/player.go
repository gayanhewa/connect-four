package players

import (
	"github.com/gayanhewa/connect-four/interactions"
)

// Player interface defines the actions at hand for a Player.
type Player interface {
	Name() string
	Move(i interactions.Interactor, availableMoves []int) (int, error)
}

package players

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gayanhewa/connect-four/interactions"
)

func TestComputerPlayerName(t *testing.T) {
	player := NewComputerPlayer()
	if player.Name() == "" {
		t.Fatal("Failed asserting that the computer player has a name")
	}
}

func TestComputerPlayerFactory(t *testing.T) {
	player := NewComputerPlayer()
	if player == nil {
		t.Fatal("Failed creating a new computer player.")
	}
}

func TestComputerPlayerMove(t *testing.T) {
	player := NewComputerPlayer()
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("answer")
	i := interactions.NewInteractor(stdout, stdin)
	availableMoves := []int{1, 4}
	move, _ := player.Move(i, availableMoves)
	if move != 1 && move != 4 {
		t.Fatalf("Failed to assert that %q is in %v", move, availableMoves)
	}
}

package players

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gayanhewa/connect-four/interactions"
)

func humanPlayerHelper() Player {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("G")
	i := interactions.NewInteractor(stdout, stdin)
	return NewHumanPlayer(i)
}
func TestNewHumanFactory(t *testing.T) {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader("answer")
	i := interactions.NewInteractor(stdout, stdin)
	player := NewHumanPlayer(i)
	if player == nil {
		t.Fatal("Failed creating a new human player.")
	}
	if player.Name() != "answer" {
		t.Fatalf("Failed asserting that %q is equal to %q", player.Name(), "answer")
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		err         string
		description string
	}{
		{"A", -1, "invalid move to column A", "Invalid move with a letter"},
		{"1", 1, "", "valid move with"},
		{"7", -1, "invalid move to column 7", "Invalid move out of bounds"},
		{"4", -1, "invalid move to column 4", "Invalid move not available"},
	}
	for _, test := range tests {
		stdout := new(bytes.Buffer)
		stdin := strings.NewReader(test.input)
		i := interactions.NewInteractor(stdout, stdin)
		player := humanPlayerHelper()
		move, err := player.Move(i, []int{1, 2, 3})
		if err != nil && err.Error() != test.err {
			t.Fatalf("Failed asserting that the move played was %q", test.input)
		}
		if move != test.expected {
			t.Fatalf("Failed asserting that the move played was %q", test.input)
		}
	}
}

func TestHumanPlayerName(t *testing.T) {
	player := humanPlayerHelper()
	if player.Name() != "G" {
		t.Fatalf("Failed asserting that the name of the player is G")
	}
}

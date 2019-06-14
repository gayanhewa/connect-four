package board

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/gayanhewa/connect-four/interactions"

	"github.com/gayanhewa/connect-four/players"
)

func humanPlayerHelper(name string) players.Player {
	stdout := new(bytes.Buffer)
	stdin := strings.NewReader(name)
	i := interactions.NewInteractor(stdout, stdin)
	return players.NewHumanPlayer(i)
}

func computerPlayerHelper() players.Player {
	return players.NewComputerPlayer()
}
func TestMove(t *testing.T) {
	board := &Board{}
	player := humanPlayerHelper("R")
	for i := 0; i < 6; i++ {
		if board.Move(0, &player) == false {
			t.Fatalf("Failed asserting succesful moves for player %q", player.Name())
		}
	}
	if board.Move(0, &player) == true {
		t.Fatalf("Failed asserting succesful moves for player %q", player.Name())
	}
}

func TestPrintBoard(t *testing.T) {
	type move struct {
		name   string
		column int
	}
	tests := []struct {
		expected    string
		moves       []move
		description string
	}{
		{"======================\n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n 0  1  2  3  4  5  6 \n ======================\n", []move{}, "Empty"},
		{"======================\n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n ¤  ¤  ¤  ¤  ¤  ¤  ¤ \n Y  ¤  ¤  ¤  ¤  ¤  ¤ \n Y  ¤  ¤  ¤  ¤  ¤  ¤ \n R  R  ¤  ¤  ¤  ¤  ¤ \n 0  1  2  3  4  5  6 \n ======================\n", []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 0},
		}, "Half way played"},
	}
	for _, test := range tests {
		board := &Board{}
		stdout := new(bytes.Buffer)
		humanPlayer := humanPlayerHelper("R")
		computerPlayer := humanPlayerHelper("Y")
		for _, move := range test.moves {
			if move.name == humanPlayer.Name() {
				board.Move(move.column, &humanPlayer)
			}
			if move.name == computerPlayer.Name() {
				board.Move(move.column, &computerPlayer)
			}
		}
		board.Print(stdout)
		if stdout.String() != test.expected {
			t.Fatalf("Failed testing %q %q %q", test.description, stdout.String(), test.expected)
		}
	}
}

func TestWinners(t *testing.T) {
	type move struct {
		name   string
		column int
	}
	tests := []struct {
		expected    string
		moves       []move
		description string
	}{
		{"R", []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 0},
			move{"R", 2},
			move{"Y", 0},
			move{"R", 3},
		}, "Horizontal win for R"},
		{"Y", []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 0},
			move{"R", 2},
			move{"Y", 0},
			move{"R", 4},
			move{"Y", 0},
		}, "Vertical win for Y"},
		{"R", []move{
			move{"R", 0},
			move{"Y", 1},
			move{"R", 1},
			move{"Y", 2},
			move{"R", 2},
			move{"Y", 3},
			move{"R", 3},
			move{"Y", 3},
			move{"R", 3},
			move{"Y", 0},
			move{"R", 2},
		}, "Diagonal win for R"},
		{"Y", []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 1},
			move{"R", 2},
			move{"Y", 1},
			move{"R", 0},
			move{"Y", 2},
			move{"R", 1},
			move{"Y", 3},
		}, "Diagonal win for Y"},
		{"error scenario", []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
		}, "no winner yet"},
	}
	for _, test := range tests {
		board := &Board{}
		humanPlayer := humanPlayerHelper("R")
		computerPlayer := humanPlayerHelper("Y")
		for _, move := range test.moves {
			if move.name == humanPlayer.Name() {
				board.Move(move.column, &humanPlayer)
			}
			if move.name == computerPlayer.Name() {
				board.Move(move.column, &computerPlayer)
			}
		}
		winner, err := board.Winner()
		if err != nil && err.Error() != test.description {
			t.Fatalf("Failed testing %q %q %q", test.description, err.Error(), test.expected)
		}
		if err == nil && winner.Name() != test.expected {
			t.Fatalf("Failed testing %q %q %q", test.description, winner.Name(), test.expected)
		}
	}
}
func TestAvailableMoves(t *testing.T) {
	type move struct {
		name   string
		column int
	}
	tests := []struct {
		expected    []int
		moves       []move
		description string
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6}, []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 0},
			move{"R", 2},
			move{"Y", 0},
			move{"R", 3},
		}, "All columns available for drop"},
		{[]int{1, 2, 3, 4, 5, 6}, []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
		}, "0th column not available for drop"},
	}
	for _, test := range tests {
		board := &Board{}
		humanPlayer := humanPlayerHelper("R")
		computerPlayer := humanPlayerHelper("Y")
		for _, move := range test.moves {
			if move.name == humanPlayer.Name() {
				board.Move(move.column, &humanPlayer)
			}
			if move.name == computerPlayer.Name() {
				board.Move(move.column, &computerPlayer)
			}
		}
		availableMoves := board.AvailableMoves()

		if !reflect.DeepEqual(availableMoves, test.expected) {
			t.Fatalf("Failed testing %q %q %q", test.description, availableMoves, test.expected)
		}
	}
}

func TestIsValidMove(t *testing.T) {
	type move struct {
		name   string
		column int
	}
	tests := []struct {
		expected    bool
		column      int
		moves       []move
		description string
	}{
		{true, 0, []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
			move{"Y", 0},
			move{"R", 2},
			move{"Y", 0},
			move{"R", 3},
		}, "All columns available for drop"},
		{false, 0, []move{
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 0},
			move{"Y", 0},
			move{"R", 1},
		}, "0th column not available for drop"},
	}
	for _, test := range tests {
		board := &Board{}
		humanPlayer := humanPlayerHelper("R")
		computerPlayer := humanPlayerHelper("Y")
		for _, move := range test.moves {
			if move.name == humanPlayer.Name() {
				board.Move(move.column, &humanPlayer)
			}
			if move.name == computerPlayer.Name() {
				board.Move(move.column, &computerPlayer)
			}
		}
		t.Logf("%v", board.moves)
		if board.IsValidMove(test.column) != test.expected {
			t.Fatalf("Failed %q", test.description)
		}
	}
}

func TestFirstMove(t *testing.T) {
	board := &Board{}
	if board.FirstMove() != true {
		t.Fatal("Failed asserting that this is the first move.")
	}
}

func TestLastMove(t *testing.T) {
	board := &Board{}
	player := humanPlayerHelper("R")
	board.Move(0, &player)
	move := board.LastMove()
	if move.Column() != 0 && move.Player() != &player {
		t.Fatal("Failed asserting that this is the first move.")
	}
}

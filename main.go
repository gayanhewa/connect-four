package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gayanhewa/connect-four/interactions"

	"github.com/gayanhewa/connect-four/board"
	"github.com/gayanhewa/connect-four/players"
)

func mainMenu() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to connect 4 !")
	fmt.Println("Who do you want to play against?")
	fmt.Println("1. Another player")
	fmt.Println("2. AI player")
	fmt.Println("3. Or just watch the dumb AI play")
	option, _ := reader.ReadString('\n')
	option = strings.TrimSuffix(option, "\n")
	optionAsInt, _ := strconv.Atoi(option)
	if optionAsInt != 1 && optionAsInt != 2 && optionAsInt != 3 {
		fmt.Println("Invalid option. Please try again.")
		mainMenu()
	}
	return optionAsInt
}

func initPlayers(interactor interactions.Interactor, option int) (players.Player, players.Player) {
	var playerOne players.Player
	var playerTwo players.Player

	if option == 1 {
		playerOne = players.NewHumanPlayer(interactor)
		playerTwo = players.NewHumanPlayer(interactor)
	}
	if option == 2 {
		playerOne = players.NewHumanPlayer(interactor)
		playerTwo = players.NewComputerPlayer()
	}
	if option == 3 {
		playerOne = players.NewComputerPlayer()
		playerTwo = players.NewComputerPlayer()
	}
	return playerOne, playerTwo
}
func main() {
	interactor := interactions.NewInteractor(os.Stdout, os.Stdin)
	playerOne, playerTwo := initPlayers(interactor, mainMenu())
	b := &board.Board{}
	for len(b.AvailableMoves()) > 0 {
		b.Print(os.Stdout)
		if !b.FirstMove() {
			m := b.LastMove()
			if m.Player() != &playerTwo {
				move, err := playerTwo.Move(interactor, b.AvailableMoves())
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				b.Move(move, &playerTwo)
			}
			if m.Player() != &playerOne {
				move, err := playerOne.Move(interactor, b.AvailableMoves())
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				b.Move(move, &playerOne)
			}
			winner, err := b.Winner()
			if err == nil {
				fmt.Printf("Game over. %s won.", winner.Name())
				b.Print(os.Stdin)
				return
			}
			continue
		}
		// Randomize the first move.
		if rand.Intn(10) < 5 {
			move, err := playerOne.Move(interactor, b.AvailableMoves())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			b.Move(move, &playerOne)
		} else {
			move, err := playerTwo.Move(interactor, b.AvailableMoves())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			b.Move(move, &playerTwo)
		}
	}
	if len(b.AvailableMoves()) < 1 {
		// Draw.
		fmt.Println("Game Over.")
		b.Print(os.Stdin)
	}
}

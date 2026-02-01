package main

import (
	"fmt"
	. "github.com/avantikasparihar/low-level-design-problems/connect-four/internal"
	"log"
	"os"
)

/*
Entities:
	Game:
		- Id: String
		- Player1: Player
		- Player2: Player
		- DefaultBoard: DefaultBoard
		- State: Enum("ONGOING", "DRAW", "PLAYER1_WON", "PLAYER2_WON")

	DefaultBoard: Array[7][6]

	Player:
		- Id: String
		- Name: String

Interfaces:
	BoardManager:
		- DropBall(column int)
		- CheckOutcome()
		- GetBoard()
*/

var game *Game

func main() {
	playerName1 := "Alice"
	playerName2 := "Bob"

	game = NewGame(playerName1, playerName2)
	move(1, 1)
	move(2, 2)
	move(1, 1)
	move(2, 3)
	move(1, 1)
	move(2, 4)
	move(1, 2)
	move(2, 5)
}

func move(player, col int) {
	state, mssg, err := game.Move(player, col)
	if err != nil {
		log.Fatalln(err)
	}
	game.Display()
	fmt.Println(mssg)
	switch state {
	case GameStateDraw, GameStatePlayer1Won, GameStatePlayer2Won:
		os.Exit(0)
	}
}

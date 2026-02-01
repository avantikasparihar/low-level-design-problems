package internal

import (
	"errors"
	"fmt"
)

const (
	GameStateOngoing    GameState = "ONGOING"
	GameStateDraw       GameState = "DRAW"
	GameStatePlayer1Won GameState = "PLAYER1_WON"
	GameStatePlayer2Won GameState = "PLAYER2_WON"
)

type Game struct {
	Id      string
	Player1 *Player
	Player2 *Player
	Board   BoardManager
	State   GameState
}

type GameState string

func NewGame(playerName1, playerName2 string) *Game {
	return &Game{
		Id:      "<generated-uid>",
		Player1: NewPlayer(1, playerName1),
		Player2: NewPlayer(2, playerName2),
		Board:   NewDefaultBoard(),
		State:   GameStateOngoing,
	}
}

func (g *Game) Move(playerId int, column int) (GameState, string, error) {
	// impl logic to orchestrate the game
	err := g.Board.DropBall(column)
	if err != nil {
		return "", "", err
	}

	outcome := g.Board.CheckOutcome()

	switch outcome {
	case OutcomeTypeHorizontal1, OutcomeTypeVertical1, OutcomeTypeDiagonal1:
		return GameStatePlayer1Won, "Game Over, Player 1 won.", nil
	case OutcomeTypeHorizontal2, OutcomeTypeVertical2, OutcomeTypeDiagonal2:
		return GameStatePlayer2Won, "Game Over, Player 2 won.", nil
	case OutcomeTypeDraw:
		return GameStateDraw, "Game Over, draw.", nil
	case OutcomeTypeUnknown:
		switch playerId {
		case 1:
			return GameStateOngoing, "Player 2's turn...", nil
		case 2:
			return GameStateOngoing, "Player 1's turn...", nil
		}
	default:
		return "", "", errors.New("unknown outcome")
	}

	return "", "", errors.New("failed")
}

func (g *Game) Display() {
	mat := g.Board.GetBoard()
	fmt.Println("Board: ")
	for i := range mat {
		for _, v := range mat[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}

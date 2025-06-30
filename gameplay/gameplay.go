// Package gameplay manages the current game state and turn logic.
package gameplay

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

// currentGame stores the state of the running game.
var currentGame types.Game

// StartNewGame initializes an empty game and sets the first active player.
func StartNewGame() {
	currentGame = types.Game{TurnNum: 1, SubTurnNum: 1, Turns: []types.Turn{}}
	players.SetFirstPlayer()
	currentGame.Turns = append(currentGame.Turns, types.Turn{Player: players.GetActivePlayer(), SubTurns: []types.SubTurn{{}}})
}

// GetTurnNum returns the current turn number.
func GetTurnNum() int {
	return currentGame.TurnNum
}

// NextTurn advances the game to the next turn and switches the active player.
func NextTurn() {
	currentGame.TurnNum++
	currentGame.SubTurnNum = 1
	players.NextPlayer()
	currentGame.Turns = append(currentGame.Turns, types.Turn{Player: players.GetActivePlayer(), SubTurns: []types.SubTurn{{}}})
}

// GetSubTurnNum returns the current subturn number.
func GetSubTurnNum() int {
	return currentGame.SubTurnNum
}

// NextSubTurn increments the subturn counter.
func NextSubTurn() {
	currentGame.SubTurnNum++
	currentGame.Turns[GetTurnNum()-1].SubTurns = append(currentGame.Turns[GetTurnNum()-1].SubTurns, types.SubTurn{})
}

// ActionDone records a single tile change that occurred during the current subturn.
func ActionDone(x, y, oldVal, newVal int) {
	sub := [4]int{x, y, oldVal, newVal}
	currentGame.Turns[GetTurnNum()-1].SubTurns[GetSubTurnNum()-1] = append(currentGame.Turns[GetTurnNum()-1].SubTurns[GetSubTurnNum()-1], sub)
}

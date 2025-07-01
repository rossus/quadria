// Package gameplay manages the current game state and turn logic.
package gameplay

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

// Game represents a single play session.
// It stores every turn and subturn that has occurred.
type Game struct {
	turnNum, subTurnNum int
	turns               []types.Turn
	players             *players.Players
}

// InitializeNewGame initializes an empty game and sets the first active player.
func InitializeNewGame(players *players.Players) *Game {
	currentGame := Game{turnNum: 1, subTurnNum: 1, turns: []types.Turn{}, players: players}
	currentGame.players.SetFirstPlayer()
	currentGame.turns = append(currentGame.turns, types.Turn{Player: currentGame.players.GetActivePlayer(), SubTurns: []types.SubTurn{{}}})
	return &currentGame
}

// GetTurnNum returns the current turn number.
func (g *Game) GetTurnNum() int {
	return g.turnNum
}

// NextTurn advances the game to the next turn and switches the active player.
func (g *Game) NextTurn() {
	g.turnNum++
	g.subTurnNum = 1
	g.players.NextPlayer()
	g.turns = append(g.turns, types.Turn{Player: g.players.GetActivePlayer(), SubTurns: []types.SubTurn{{}}})
}

// GetSubTurnNum returns the current subturn number.
func (g *Game) GetSubTurnNum() int {
	return g.subTurnNum
}

// NextSubTurn increments the subturn counter.
func (g *Game) NextSubTurn() {
	g.subTurnNum++
	g.turns[g.GetTurnNum()-1].SubTurns = append(g.turns[g.GetTurnNum()-1].SubTurns, types.SubTurn{})
}

// ActionDone records a single tile change that occurred during the current subturn.
func (g *Game) ActionDone(x, y, oldVal, newVal int) {
	sub := [4]int{x, y, oldVal, newVal}
	g.turns[g.GetTurnNum()-1].SubTurns[g.GetSubTurnNum()-1] = append(g.turns[g.GetTurnNum()-1].SubTurns[g.GetSubTurnNum()-1], sub)
}

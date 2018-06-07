package gameplay

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

var currentGame types.Game

func StartNewGame() {
	currentGame = types.Game{1, 1, []types.Turn{}}
	players.SetFirstPlayer()
	currentGame.Turns = append(currentGame.Turns, types.Turn{players.GetActivePlayer(), []types.SubTurn{{}}})
}

func GetTurnNum() int {
	return currentGame.TurnNum
}

func NextTurn() {
	currentGame.TurnNum++
	currentGame.SubTurnNum = 1
	players.NextPlayer()
	currentGame.Turns = append(currentGame.Turns, types.Turn{players.GetActivePlayer(), []types.SubTurn{{}}})
}

func GetSubTurnNum() int {
	return currentGame.SubTurnNum
}

func NextSubTurn() {
	currentGame.SubTurnNum++
	currentGame.Turns[GetTurnNum()-1].SubTurns = append(currentGame.Turns[GetTurnNum()-1].SubTurns, types.SubTurn{})
}

func ActionDone(x, y, oldVal, newVal int) {
	sub := [4]int{x, y, oldVal, newVal}
	currentGame.Turns[GetTurnNum()-1].SubTurns[GetSubTurnNum()-1] = append(currentGame.Turns[GetTurnNum()-1].SubTurns[GetSubTurnNum()-1], sub)
}

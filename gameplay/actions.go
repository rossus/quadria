package gameplay

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/board"
	"github.com/rossus/quadria/players"
)

func addToNext(nextActions []types.NextAction, checked [][3]int) []types.NextAction {
	for i := 0; i < len(checked); i++ {
		for j := 0; j < len(nextActions); j++ {
			nx := nextActions[j].X
			ny := nextActions[j].Y
			cx := checked[i][0]
			cy := checked[i][1]
			if (nx == cx) && (ny == cy) {
				nextActions[j].Amount = nextActions[j].Amount + checked[i][2]
				i++
				break
			}
		}
		nextActions = append(nextActions, types.NextAction{checked[i][0], checked[i][1], checked[i][2]})
	}
	return nextActions
}

func checkOverlap(x, y int) [][3]int {
	var futureActions [][3]int
	tile := board.GetTile(x, y)
	if tile.Value <= tile.Neighbours {
		return nil
	} else {
		grade := tile.Value / tile.Neighbours
		loose := -tile.Neighbours * grade
		futureActions = append(futureActions, [3]int{x, y, loose})
		if x != 0 {
			futureActions = append(futureActions, [3]int{x - 1, y, grade})
		}
		if x != len(board.GetBoard().Tiles) {
			futureActions = append(futureActions, [3]int{x + 1, y, grade})
		}
		if y != 0 {
			futureActions = append(futureActions, [3]int{x, y - 1, grade})
		}
		if y != len(board.GetBoard().Tiles) {
			futureActions = append(futureActions, [3]int{x, y + 1, grade})
		}
	}
	return futureActions
}

func goSub(currentActions []types.NextAction) []types.NextAction {
	nextActions := make([]types.NextAction, 0)
	for i := 0; i < len(currentActions); i++ {
		//Action execution
		oldAmount := board.GetTile(currentActions[i].X, currentActions[i].Y).Value
		board.ChangeTileState(currentActions[i].X, currentActions[i].Y, oldAmount+currentActions[i].Amount)
		ActionDone(currentActions[i].X, currentActions[i].Y, oldAmount, oldAmount+currentActions[i].Amount)

		//Plan actions for the next subturn
		nextActions = addToNext(nextActions, checkOverlap(currentActions[i].X, currentActions[i].Y))
	}

	return nextActions
}

func Go(x, y int) {
	if (board.GetTile(x, y).Player == players.GetActivePlayer() || board.GetTile(x, y).Player == players.GetBlankPlayer()) {
		nextActions := []types.NextAction{{x, y, 1}}
		for {
			nextActions = goSub(nextActions)
			if len(nextActions) == 0 {
				NextTurn()
				break
			}

			NextSubTurn()
		}
	}
}

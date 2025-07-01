package session

import (
	"github.com/rossus/quadria/board"
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/gameplay"
	"github.com/rossus/quadria/players"
)

type Session struct {
	Players *players.Players
	Board   *board.Board
	Game    *gameplay.Game
}

func InitializeNewSession(players *players.Players, board *board.Board, game *gameplay.Game) *Session {
	return &Session{
		Players: players,
		Board:   board,
		Game:    game,
	}
}

// Go performs a player's move on the given tile and runs any resulting chain reactions.
// It returns true if the move ended the game.
func (s *Session) Go(x, y int) bool {
	if s.Board.GetTile(x, y).Player == s.Players.GetActivePlayer() || s.Board.GetTile(x, y).Player == s.Players.GetBlankPlayer() {
		nextActions := []types.NextAction{{X: x, Y: y, Amount: 1}}
		var chk bool
		for {
			nextActions, chk = s.goSub(nextActions)
			if chk {
				return true
			}
			if len(nextActions) == 0 {
				s.Game.NextTurn()
				break
			}

			s.Game.NextSubTurn()
		}
	}
	return false
}

// addToNext merges the checked actions with the slice of next actions,
// aggregating amounts when the same tile is encountered.
func (s *Session) addToNext(nextActions []types.NextAction, checked [][3]int) []types.NextAction {
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

		if i < len(checked) {
			nextActions = append(nextActions, types.NextAction{X: checked[i][0], Y: checked[i][1], Amount: checked[i][2]})
		}
	}
	return nextActions
}

// checkOverlap returns the list of actions triggered when a tile exceeds its neighbours.
func (s *Session) checkOverlap(x, y int) [][3]int {
	var futureActions [][3]int
	tile := s.Board.GetTile(x, y)
	if tile.Value <= tile.Neighbours {
		return nil
	} else {
		grade := tile.Value / tile.Neighbours
		loose := -tile.Neighbours * grade
		futureActions = append(futureActions, [3]int{x, y, loose})
		if x != 0 {
			futureActions = append(futureActions, [3]int{x - 1, y, grade})
		}
		if x != len(s.Board.GetTiles())-1 {
			futureActions = append(futureActions, [3]int{x + 1, y, grade})
		}
		if y != 0 {
			futureActions = append(futureActions, [3]int{x, y - 1, grade})
		}
		if y != len(s.Board.GetTiles())-1 {
			futureActions = append(futureActions, [3]int{x, y + 1, grade})
		}
	}
	return futureActions
}

// goSub performs one subturn and returns actions for the next subturn.
// The bool result is true when the game is finished.
func (s *Session) goSub(currentActions []types.NextAction) ([]types.NextAction, bool) {
	nextActions := make([]types.NextAction, 0)
	for i := 0; i < len(currentActions); i++ {
		//Action execution
		oldAmount := s.Board.GetTile(currentActions[i].X, currentActions[i].Y).Value
		s.Board.ChangeTileState(currentActions[i].X, currentActions[i].Y, oldAmount+currentActions[i].Amount)
		s.Game.ActionDone(currentActions[i].X, currentActions[i].Y, oldAmount, oldAmount+currentActions[i].Amount)

		//Check if one player owns an entire board, if so end the game
		if s.Board.CheckDomination() {
			return nil, true
		}

		//Plan actions for the next subturn
		nextActions = s.addToNext(nextActions, s.checkOverlap(currentActions[i].X, currentActions[i].Y))
	}

	return nextActions, false
}

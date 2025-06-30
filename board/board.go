// Package board contains helpers to manage the game board state.
package board

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

// board holds the current state of the game board.
var board types.Board

// InitNewBoard creates a new square board of the provided size and stores it as the active board.
func InitNewBoard(size int) {
	var newBoard types.Board
	for i := 0; i < size; i++ {
		var row []types.Tile
		for j := 0; j < size; j++ {
			var tile types.Tile
			if (i == 0 || i == size-1) && (j == 0 || j == size-1) {
				tile = types.Tile{Value: 1, Neighbours: 2, Player: players.GetBlankPlayer()}
			} else if (i == 0 || i == size-1) || (j == 0 || j == size-1) {
				tile = types.Tile{Value: 1, Neighbours: 3, Player: players.GetBlankPlayer()}
			} else {
				tile = types.Tile{Value: 1, Neighbours: 4, Player: players.GetBlankPlayer()}
			}
			row = append(row, tile)
		}
		newBoard.Tiles = append(newBoard.Tiles, row)
	}
	board = newBoard
}

// GetBoard returns the current board state.
func GetBoard() types.Board {
	return board
}

// GetTile returns the tile located at the given coordinates.
func GetTile(x, y int) types.Tile {
	return board.Tiles[y][x]
}

// ChangeTileState sets the value of a tile and assigns it to the active player.
func ChangeTileState(x, y, newVal int) {
	board.Tiles[y][x].Value = newVal
	board.Tiles[y][x].Player = players.GetActivePlayer()
}

// CheckDomination reports whether a single player controls the entire board.
func CheckDomination() bool {
	if *board.Tiles[0][0].Player != *players.GetBlankPlayer() {
		player := *board.Tiles[0][0].Player
		for i := 0; i < len(board.Tiles); i++ {
			for j := 0; j < len(board.Tiles[i]); j++ {
				if *board.Tiles[i][j].Player != player {
					return false
				}
			}
		}
		return true
	}
	return false
}

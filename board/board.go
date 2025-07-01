// Package board contains helpers to manage the game board state.
package board

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

// Board holds a 2D slice of Tiles that make up the play field.
type Board struct {
	tiles   [][]types.Tile
	players *players.Players
}

// InitNewBoard creates a new square board of the provided size and stores it as the active board.
func InitNewBoard(size int, players *players.Players) *Board {
	newBoard := Board{
		players: players,
	}

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
		newBoard.tiles = append(newBoard.tiles, row)
	}
	return &newBoard
}

// GetBoard returns the current board state.
func (b *Board) GetBoard() Board {
	return *b
}

// GetTiles exposes the current board tile matrix.
func (b *Board) GetTiles() [][]types.Tile {
	return b.tiles
}

// GetTile returns the tile located at the given coordinates.
func (b *Board) GetTile(x, y int) types.Tile {
	return b.tiles[y][x]
}

// ChangeTileState sets the value of a tile and assigns it to the active player.
func (b *Board) ChangeTileState(x, y, newVal int) {
	b.tiles[y][x].Value = newVal
	b.tiles[y][x].Player = b.players.GetActivePlayer()
}

// CheckDomination reports whether a single player controls the entire board.
func (b *Board) CheckDomination() bool {
	if *b.tiles[0][0].Player != *b.players.GetBlankPlayer() {
		player := *b.tiles[0][0].Player
		for i := 0; i < len(b.tiles); i++ {
			for j := 0; j < len(b.tiles[i]); j++ {
				if *b.tiles[i][j].Player != player {
					return false
				}
			}
		}
		return true
	}
	return false
}

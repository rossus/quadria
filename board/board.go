package board

import (
	"github.com/rossus/quadria/common/types"
	"github.com/rossus/quadria/players"
)

var board types.Board

func InitNewBoard(size int) {
	var newBoard types.Board
	for i := 0; i < size; i++ {
		var row []types.Tile
		for j := 0; j < size; j++ {
			var tile types.Tile
			if (i == 0 || i == size-1) && (j == 0 || j == size-1) {
				tile = types.Tile{1, 2, players.GetBlankPlayer()}
			} else if (i == 0 || i == size-1) || (j == 0 || j == size-1) {
				tile = types.Tile{1, 3, players.GetBlankPlayer()}
			} else {
				tile = types.Tile{1, 4, players.GetBlankPlayer()}
			}
			row = append(row, tile)
		}
		newBoard.Tiles = append(newBoard.Tiles, row)
	}
	board = newBoard
}

func GetBoard() types.Board {
	return board
}

func GetTile(x, y int) types.Tile {
	return board.Tiles[y][x]
}

func ChangeTileState(x, y, newVal int) {
	board.Tiles[y][x].Value = newVal
	board.Tiles[y][x].Player = players.GetActivePlayer()
}

func CheckDomination() bool {
	if *board.Tiles[0][0].Player != *players.GetBlankPlayer() {
		player:=*board.Tiles[0][0].Player
		for i:=0; i<len(board.Tiles); i++ {
			for j:=0; j<len(board.Tiles[i]); j++ {
				if *board.Tiles[i][j].Player != player {
					return false
				}
			}
		}
		return true
	}
	return false
}

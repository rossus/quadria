// Package types defines shared structures used throughout Quadria.
package types

// Tile represents a single square on the board.
// Value is the current number stored in the tile.
// Neighbours defines how many adjacent tiles it has.
// Player points to the owner of the tile.
type Tile struct {
	Value      int
	Neighbours int
	Player     *Player
}

// Board holds a 2D slice of Tiles that make up the play field.
type Board struct {
	Tiles [][]Tile
}

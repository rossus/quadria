package types

type Tile struct {
	Value int
	Neighbours int
	Player *Player
}

type Board struct {
	Tiles [][]Tile
}
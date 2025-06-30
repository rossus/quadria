// Package types defines shared structures used throughout Quadria.
package types

// Game represents a single play session.
// It stores every turn and subturn that has occurred.
type Game struct {
	TurnNum, SubTurnNum int
	Turns               []Turn
}

// Turn stores all subturns performed by one player during a turn.
type Turn struct {
	Player   *Player
	SubTurns []SubTurn
}

// SubTurn is a slice of tile changes. Each tile change is an array of four values:
// [x coordinate, y coordinate, initial tile value, new tile value].
type SubTurn [][4]int

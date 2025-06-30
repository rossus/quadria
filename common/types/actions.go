// Package types defines shared structures used throughout Quadria.
package types

// NextAction describes an amount of value to add to a tile at (X,Y).
type NextAction struct {
	X, Y, Amount int
}

package types

type Game struct {
	TurnNum, SubTurnNum int
	Turns []Turn
}

type Turn struct {
	Player *Player
	SubTurns []SubTurn
}

//Subturn is a slice of tile changes. Each tile change is an array of 4 parameters:
//[x coordinate, y coordinate, initial tile value, new tile value]
type SubTurn [][4]int
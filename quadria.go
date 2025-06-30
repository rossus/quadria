// Quadria is a simple console based implementation of a chain reaction style game.
package main

import (
	"fmt"
	"github.com/rossus/quadria/board"
	"github.com/rossus/quadria/controller"
	"github.com/rossus/quadria/gameplay"
	"github.com/rossus/quadria/players"
)

// main initializes players, creates the board and starts the game loop.
func main() {
	var name string
	fmt.Print("Enter name of the blue player: ")
	fmt.Scanln(&name)
	players.InitPlayer(name, "blue")
	fmt.Print("Enter name of the red player: ")
	fmt.Scanln(&name)
	players.InitPlayer(name, "red")

	var size int
	fmt.Print("Enter size: ")
	fmt.Scanln(&size)
	board.InitNewBoard(size)

	gameplay.StartNewGame()
	controller.CHRun()

}

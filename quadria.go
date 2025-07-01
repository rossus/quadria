// Quadria is a simple console based implementation of a chain reaction style game.
package main

import (
	"fmt"

	"github.com/rossus/quadria/board"
	"github.com/rossus/quadria/controller"
	"github.com/rossus/quadria/gameplay"
	"github.com/rossus/quadria/players"
	"github.com/rossus/quadria/session"
)

// main initializes players, creates the board and starts the game loop.
func main() {
	chPlayers := players.InitPlayers()
	var name string
	fmt.Print("Enter name of the blue player: ")
	fmt.Scanln(&name)
	chPlayers.AddPlayer(name, "blue")
	fmt.Print("Enter name of the red player: ")
	fmt.Scanln(&name)
	chPlayers.AddPlayer(name, "red")

	var size int
	fmt.Print("Enter size: ")
	fmt.Scanln(&size)
	chBoard := board.InitNewBoard(size, chPlayers)

	chGame := gameplay.InitializeNewGame(chPlayers)

	chSession := session.InitializeNewSession(chPlayers, chBoard, chGame)
	chController := controller.InitializeConsoleHotseatController(chSession)
	chController.Run()
}

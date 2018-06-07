package main

import (
	"fmt"
	"github.com/rossus/quadra/players"
	"github.com/rossus/quadra/board"
	"github.com/rossus/quadra/gameplay"
	"github.com/rossus/quadra/controller"
)

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

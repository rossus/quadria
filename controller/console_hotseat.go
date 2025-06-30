// Package controller implements a simple console hotseat interface.
package controller

import (
	"bufio"
	"fmt"
	"github.com/kortschak/ct"
	"github.com/rossus/quadria/board"
	"github.com/rossus/quadria/gameplay"
	"github.com/rossus/quadria/players"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/buger/goterm"
)

// TODO: Try to use github.com/buger/goterm for fix
// clear attempts to clear the terminal screen across platforms.
func clear() {
	var c *exec.Cmd
	var doClear = true

	switch runtime.GOOS {
	case "darwin":
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		doClear = false
	}
	if doClear {
		c.Stdout = os.Stdout
		c.Run()
	}
}

// getColor converts a player color name into ct color attributes.
func getColor(color string) (ct.Color, ct.Mode) {
	switch color {
	case "gray":
		return ct.White, ct.Bold
	case "blue":
		return ct.Blue, ct.Fg(ct.BoldYellow)
	case "red":
		return ct.Red, ct.Bold
	default:
		return ct.Green, ct.Bold
	}
}

// drawBoard prints the current game board to the terminal using color.
func drawBoard() {
	var CPBg, CPTxt = getColor(players.GetActivePlayer().Color)
	var currentPlayer = (ct.Bg(CPBg) | CPTxt).Paint
	var currBoard = board.GetBoard()

	fmt.Println(currentPlayer("Turn ", gameplay.GetTurnNum()))
	fmt.Println()

	for i := 0; i < len(currBoard.Tiles); i++ {
		fmt.Println()
		for j := 0; j < len(currBoard.Tiles[i]); j++ {
			var TOBg, TOTxt = getColor(currBoard.Tiles[i][j].Player.Color)
			var tileOwner = (ct.Bg(TOBg) | TOTxt).Paint
			fmt.Print(tileOwner(currBoard.Tiles[i][j].Value))
		}
	}
}

// CHRun runs the console hotseat game loop.
func CHRun() {
	fmt.Println("Welcome to the Quadria console hotseat game! There are two players here: blue (1) and red (2).")
	fmt.Println("It's turn 1 now. It is blue's turn.")
	for {
		drawBoard()
		fmt.Println()
		fmt.Println("What would you do? Type 'help' to see commands.")
		var command string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command = scanner.Text()
		}
		cmd := strings.Split(command, " ")
		if cmd[0] == "exit" {
			fmt.Println("Bye!")
			break
		} else if cmd[0] == "go" {
			if len(cmd) >= 3 {
				x, err := strconv.Atoi(cmd[1])
				if err != nil {
					fmt.Println(err)
				} else if y, err := strconv.Atoi(cmd[2]); err != nil {
					fmt.Println(err)
				} else if x >= 0 && x < len(board.GetBoard().Tiles) && y >= 0 && y < len(board.GetBoard().Tiles) {
					if gameplay.Go(x, y) {
						goterm.Clear()
						goterm.MoveCursor(1, 1)
						goterm.Flush()
						drawBoard()
						fmt.Println()
						fmt.Println("Player ", players.GetActivePlayer().Name, " won!")
						fmt.Print("Type anything for exit...")
						var anyKey string
						fmt.Scanln(&anyKey)
						break
					}
				}
			}
		} else if cmd[0] == "help" {
			fmt.Println("You can use these commands:")
			fmt.Println("go [x] [y]						//add 1 to the tile (x, y)")
			fmt.Println("exit							//exit game")
			fmt.Println("help							//see all commands")
		}
		goterm.Clear()
		goterm.MoveCursor(1, 1)
		goterm.Flush()
	}
}

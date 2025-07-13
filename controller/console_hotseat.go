// Package controller implements a simple console hotseat interface.
package controller

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rossus/quadria/session"

	"github.com/buger/goterm"
)

// ConsoleHotseatController provides a simple text-based hotseat interface.
type ConsoleHotseatController struct {
	session *session.Session
}

// InitializeConsoleHotseatController constructs a console controller for the provided session.
func InitializeConsoleHotseatController(session *session.Session) *ConsoleHotseatController {
	return &ConsoleHotseatController{
		session: session,
	}
}

// Run executes the console hotseat game loop.
func (chc *ConsoleHotseatController) Run() {
	fmt.Println("Welcome to the Quadria console hotseat game! There are two players here: blue (1) and red (2).")
	fmt.Println("It's turn 1 now. It is blue's turn.")
	for {
		chc.drawBoard()
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
				} else if x >= 0 && x < len(chc.session.Board.GetTiles()) && y >= 0 && y < len(chc.session.Board.GetTiles()) {
					if chc.session.Go(x, y) {
						goterm.Clear()
						goterm.MoveCursor(1, 1)
						goterm.Flush()
						chc.drawBoard()
						fmt.Println()
						fmt.Println("Player ", chc.session.Players.GetActivePlayer().Name, " won!")
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

// drawBoard prints the current game board to the terminal using color.
func (chc *ConsoleHotseatController) drawBoard() {
	var currentPlayer = getColor(chc.session.Players.GetActivePlayer().Color).SprintFunc()
	var currentBoardTiles = chc.session.Board.GetTiles()

	fmt.Println(currentPlayer("Turn ", chc.session.Game.GetTurnNum()))
	fmt.Println()

	for i := 0; i < len(currentBoardTiles); i++ {
		fmt.Println()
		for j := 0; j < len(currentBoardTiles[i]); j++ {
			var tileOwner = getColor(currentBoardTiles[i][j].Player.Color).SprintFunc()
			fmt.Print(tileOwner(currentBoardTiles[i][j].Value))
		}
	}
}

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

// getColor converts a player color name into color.
func getColor(colour string) *color.Color {
	switch colour {
	case "gray":
		return color.New(color.FgWhite, color.Bold)
	case "blue":
		return color.New(color.FgYellow, color.Bold, color.BgBlue)
	case "red":
		return color.New(color.FgYellow, color.Bold, color.BgRed)
	default:
		return color.New(color.FgGreen, color.Bold)
	}
}

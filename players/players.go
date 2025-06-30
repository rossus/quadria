// Package players provides helper functions for player management.
package players

import (
	"github.com/rossus/quadria/common/types"
)

// players holds all players taking part in the game. The first element is a blank player used for unowned tiles.
var players = []types.Player{{Name: "", Color: "gray"}}

// activePlayer references the player whose turn is currently in progress.
var activePlayer *types.Player

// InitPlayer registers a new player with the given name and color.
func InitPlayer(name, color string) {
	players = append(players, types.Player{Name: name, Color: color})
}

// GetBlankPlayer returns the placeholder player representing no owner.
func GetBlankPlayer() *types.Player {
	return &players[0]
}

// GetActivePlayer returns the player whose turn is active.
func GetActivePlayer() *types.Player {
	return activePlayer
}

// SetFirstPlayer sets the first real player as active.
func SetFirstPlayer() {
	activePlayer = &players[1]
}

// NextPlayer advances the active player pointer to the next registered player.
func NextPlayer() {
	for i := 0; i < len(players); i++ {
		if *activePlayer == players[i] {
			if i == len(players)-1 {
				SetFirstPlayer()
				break
			} else {
				activePlayer = &players[i+1]
				break
			}
		}
	}
}

// Package players provides helper functions for player management.
package players

import (
	"github.com/rossus/quadria/common/types"
)

type Players struct {
	// players holds all players taking part in the game. The first element is a blank player used for unowned tiles.
	players []types.Player
	// activePlayer references the player whose turn is currently in progress.
	activePlayer *types.Player
}

func InitPlayers() *Players {
	return &Players{
		players: []types.Player{{Name: "", Color: "gray"}},
	}
}

// AddPlayer registers a new player with the given name and color.
func (p *Players) AddPlayer(name, color string) {
	p.players = append(p.players, types.Player{Name: name, Color: color})
}

// GetBlankPlayer returns the placeholder player representing no owner.
func (p *Players) GetBlankPlayer() *types.Player {
	return &p.players[0]
}

// GetActivePlayer returns the player whose turn is active.
func (p *Players) GetActivePlayer() *types.Player {
	return p.activePlayer
}

// SetFirstPlayer sets the first real player as active.
func (p *Players) SetFirstPlayer() {
	p.activePlayer = &p.players[1]
}

// NextPlayer advances the active player pointer to the next registered player.
func (p *Players) NextPlayer() {
	for i := 0; i < len(p.players); i++ {
		if *p.activePlayer == p.players[i] {
			if i == len(p.players)-1 {
				p.SetFirstPlayer()
				break
			} else {
				p.activePlayer = &p.players[i+1]
				break
			}
		}
	}
}

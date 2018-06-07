package players

import (
	"github.com/rossus/quadra/common/types"
)

var players = []types.Player{{"", "gray"}}
var activePlayer *types.Player

func InitPlayer(name, color string) {
	players = append(players, types.Player{name, color})
}

func GetBlankPlayer() *types.Player {
	return &players[0]
}

func GetActivePlayer() *types.Player {
	return activePlayer
}

func SetFirstPlayer() {
	activePlayer = &players[1]
}

func NextPlayer() {
	for i := 0; i < len(players); i++ {
		if *activePlayer == players[i] {
			if i==len(players)-1 {
				SetFirstPlayer()
				break
			} else {
				activePlayer = &players[i+1]
				break
			}
		}
	}
}

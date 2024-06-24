package game

import (
	"github.com/punkplod23/battle-system-go/internal/app/gameengine"
)

func Init() {

	game := gameengine.NewGameEngine()
	game.StartGame()

}

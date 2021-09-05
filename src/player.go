package src

import (
	"alberto/src/res"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	running [3]ebiten.Image
	shoot   ebiten.Image
	jump    ebiten.Image

	state int
}

var player Player

func init() {
	player = Player{
		shoot: *res.GetSqr(4, 4, "player"),
	}
	copy(player.running[:], res.GetLine("player", 3, 0)[:3])
}

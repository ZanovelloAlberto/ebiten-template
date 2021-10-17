package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	time = 40
)

var Ploading = &Loading{
	presence: true,
	current:  0,
}

type Loading struct {
	current  int
	presence bool
}

func (u *Loading) Update() {
	u.current += 1
	if time < u.current {
		u.presence = false
	}
	// println(u.current)
}

func (u Loading) Draw(screen *ebiten.Image) {
	if u.presence {

		ebitenutil.DebugPrint(screen, "loading")
	}
}

package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	time = 10
)

var Ploading Loading = Loading{
	presence: true,
}

type Loading struct {
	current  int
	presence bool
}

func (u Loading) Update() {
	u.current++
	if time < u.current {
		u.presence = false
	}
}

func (u Loading) Draw(screen *ebiten.Image) {
	if u.presence {

		ebitenutil.DebugPrint(screen, "loading")
	}
}

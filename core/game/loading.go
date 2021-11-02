package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	time = 40
)

const (
	screenWidth  = 400
	screenHeight = 400
)

var Vloading = &Loading{
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

		ebitenutil.DebugPrintAt(screen, "randome", 0, screenHeight/2)
		// ebitenutil.DebugPrint(screen, "loading")
		// ebitenutil.DebugPrint(screen, "culo")
	}
}

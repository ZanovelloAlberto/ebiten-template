package core

import (
	"egame/src/game/images"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	positions [12]ebiten.Image

	running bool

	direction Direction

	state State

	speed int
}

type State uint8
type Direction uint8

const (
	right State = iota // 0
	idle               // 1
	left               // 2
)

const (
	Dfront Direction = iota
	Dleft
	Dright
	Dback
)

var ()

var MainChar = Player{
	positions: images.GetCharacter(0),
	state:     idle,
	direction: Dfront,
	speed:     2,
	running:   false,
}

func init() {

}

// DRAW =====================
func (p Player) Draw(screen *ebiten.Image) {
	screen.DrawImage(&MainChar.positions[0], &ebiten.DrawImageOptions{})
}

// UPDATE =====================

func (p Player) Update() {
	// inpututil.AppendPressedKeys()
}

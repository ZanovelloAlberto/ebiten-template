package sprite

import (
	// "core/game/images"

	"github.com/ZanovelloAlberto/EbitenGame/core/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	images [12]ebiten.Image

	running bool

	direction Direction

	state State

	speed int

	point Point
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

var PPlayer = Player{
	images:    assets.GetCharacter(0),
	state:     idle,
	direction: Dfront,
	speed:     2,
	running:   false,
}

func init() {

}

// DRAW =====================
func (p *Player) Draw(screen *ebiten.Image) {
	var opt = &ebiten.DrawImageOptions{}

	opt.GeoM.Scale(2, 2)
	// opt.GeoM.Translate(p)
	screen.DrawImage(&PPlayer.images[p.state], opt)
}

// UPDATE =====================

func (p *Player) Update() {
	p.state = (p.state + 1) % 3
	println(p.state)
}

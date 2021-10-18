package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	time = 40
)

func init() {
	mplusBigFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

var (
	mplusBigFont font.Face
)
var Ploading = &Loading{
	presence: true,
	current:  0,
}

type Loading struct {
	current  int
	presence bool
	color    color.RGBA
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

		u.color.G = 0x80 + uint8(rand.Intn(0x7f))
		u.color.B = 0x80 + uint8(rand.Intn(0x7f))
		u.color.R = 0x80 + uint8(rand.Intn(0x7f))
		u.color.A = 0xff

		var Plan9 = []color.Color{}
		a := Plan9[0]
		text.Draw(screen, "Loading", fonts.MPlus1pRegular_ttf, screen, screenWidth, screenHeight)
		ebitenutil.DebugPrint(screen, "loading")
	}
}

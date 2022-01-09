package core

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontMagico font.Face
	colore     = color.RGBA{0xed, 0xe0, 0xc8, 0xff}
)

func init() {
	const dpi = 20

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
	}
	fontMagico, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type TextField struct {
	size           image.Point
	margin         []int
	backgroudColor color.RGBA
}

func NewTextField(w, h int, backgroudColor color.RGBA) *TextField {
	b := new(TextField)
	b.size = image.Pt(w, h)
	b.margin = []int{0, 0, 0, 0}
	b.backgroudColor = backgroudColor
	return b
}

func (b *TextField) Size() (int, int) {
	return b.size.X, b.size.Y
}

func (b *TextField) Margin() []int {
	return b.margin
}

func (b *TextField) SetMargin(m []int) {
	b.margin = m
}

func (b *TextField) Draw(screen *ebiten.Image, frame image.Rectangle) {

	text.Draw(screen, "unosre", fontMagico, 100, 100, color.RGBA{0xed, 0xe0, 0xc8, 0xff})
	// shared.FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0xff})
	// shared.DrawRect(screen, frame, b.backgroudColor, 2)
	// ebitenutil.DebugPrintAt(screen, "TextField",
	// 	frame.Min.X+((frame.Dx()-36)/2), frame.Min.Y+b.size.Y/2-8)
}

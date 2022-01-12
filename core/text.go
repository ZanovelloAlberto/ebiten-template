package core

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	uno "github.com/hajimehoshi/ebiten/v2/text"
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
	Text           *string
}

func NewTextField(size int, backgroudColor color.RGBA, text *string) *TextField {
	b := new(TextField)
	textSize := uno.BoundString(fontMagico, "*b.ddTexkijjhjt")

	b.size = image.Pt(textSize.Dx(), textSize.Dy())
	b.margin = []int{0, 0, 0, 0}
	b.backgroudColor = backgroudColor
	b.Text = text

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

	bo := text.BoundString(fontMagico, *b.Text)
	text.Draw(
		screen,
		*b.Text,
		fontMagico,
		frame.Min.X+frame.Dx()/2-bo.Dx()/2,
		frame.Min.Y+frame.Dy()/2+bo.Dy()/2,
		b.backgroudColor,
	)
}

package core

import (
	"fmt"
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
	Text           string
}

func NewTextField(backgroudColor color.RGBA, text string) *TextField {
	b := new(TextField)
	b.size = image.Pt(80, 40)
	b.margin = []int{10, 10, 10, 10}
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

	value := uno.BoundString(fontMagico, fmt.Sprint(Score)).Add(image.Pt(frame.Min.X, frame.Min.Y))
	value.Max.X += frame.Dx() / 2
	value.Max.Y += frame.Dy() / 2
	FillRect(screen, frame, color.NRGBA{0xa3, 0x49, 0xa4, 0xff})
	DrawRect(screen, frame, color.NRGBA{0xa3, 0x49, 0xa4, 0xff}, 2)

	bo := text.BoundString(fontMagico, b.Text)
	text.Draw(
		screen,
		b.Text,
		fontMagico,
		frame.Min.X+frame.Dx()/2-bo.Dx()/2,
		frame.Min.Y+frame.Dy()/2+bo.Dy()/2,
		color.RGBA{0xee, 0xe4, 0xda, 0xff},
	)

}

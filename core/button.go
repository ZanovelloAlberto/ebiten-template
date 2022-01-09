package core

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Button struct {
	size       image.Point
	isPressing bool
	margin     []int
	OnClick    func()
	Text       string
}

func NewButton(w, h int, onClick func(), text string) *Button {
	b := new(Button)
	b.size = image.Pt(w, h)
	b.margin = []int{0, 0, 0, 0}
	b.OnClick = onClick
	b.Text = text
	return b
}

func (b *Button) Size() (int, int) {
	return b.size.X, b.size.Y
}

func (b *Button) Margin() []int {
	return b.margin
}

func (b *Button) SetMargin(m []int) {
	b.margin = m
}

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.isPressing = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.isPressing = false
	if isCancel {
		// println("The click is cancelled!")
	} else {
		// println("clicked!")
		b.OnClick()
	}
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if b.isPressing {
		FillRect(screen, frame, color.RGBA{0xaa, 0, 0, 0xff})
	} else {
		FillRect(screen, frame, color.RGBA{0, 0xaa, 0, 0xff})
	}
	DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
	text.Draw(
		screen,
		b.Text,
		fontMagico,
		frame.Min.X+frame.Dx()/2,
		frame.Min.Y+frame.Dy()/2,
		color.RGBA{0xed, 0xe0, 0xc8, 0xff},
	)

}

package core

import (
	"image"

	"github.com/ZanovelloAlberto/EbitenGame/core/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Picture struct {
	img  *ebiten.Image
	size image.Point
}

func NewPicture(img []byte, w int, h int) (res *Picture) {
	res = &Picture{}
	res.size = image.Pt(w, h)
	res.img = assets.GetImage(&img)

	return res

}
func (p *Picture) Size() (int, int) {
	return p.size.X, p.size.Y
}

func (pic *Picture) Draw(screen *ebiten.Image, frame image.Rectangle) {

	d := ebiten.DrawImageOptions{}
	w, h := pic.img.Size()
	uno, due := float64(frame.Max.X-frame.Min.X)/float64(w), float64(frame.Max.Y-frame.Min.Y)/float64(h)
	a, b := float64(frame.Min.X), float64(frame.Min.Y)
	d.GeoM.Scale(uno, due)
	d.GeoM.Translate(a, b)

	screen.DrawImage(pic.img, &d)
}

package twenty48

import (
	"fmt"
	"image"

	"github.com/ZanovelloAlberto/EbitenGame/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Picture struct {
	img  *ebiten.Image
	size image.Point
}

func NewPicture(img []byte, w int, h int) (res *Picture) {
	res = &Picture{}
	res.size = image.Pt(w, h)
	res.img = core.GetImage(&img)

	println(res.img)
	return res

}
func (p *Picture) Size() (int, int) {
	return p.size.X, p.size.Y
}

func (pic *Picture) Draw(screen *ebiten.Image, frame image.Rectangle) {

	// printing the values of x and y
	fmt.Printf("%f \n", float64(frame.Max.X-frame.Min.X)/float64(pic.size.X))
	// shared.FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0xff})
	d := ebiten.DrawImageOptions{}
	// println(pic.img.Size())
	// println(frame.Max.X-frame.Min.X, frame.Max.Y-frame.Min.Y)
	// println(float64(frame.Min.X), float64(frame.Min.Y))
	// println(frame.Min.X, frame.Min.Y)
	w, h := pic.img.Size()
	uno, due := float64(frame.Max.X-frame.Min.X)/float64(w), float64(frame.Max.Y-frame.Min.Y)/float64(h)
	a, b := float64(frame.Min.X), float64(frame.Min.Y)
	d.GeoM.Scale(uno, due)
	d.GeoM.Translate(a, b)

	screen.DrawImage(pic.img, &d)
}

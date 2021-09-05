package res

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	size = 16
)

var (
	Player []ebiten.Image
)

func init() {

	// img, _, _ := ebitenutil.NewImageFromFile("res/graphics/player.png")
	// Player = pack(*img, 3)

}

func getPng(s string) (r *ebiten.Image) {
	r, _, _ = (ebitenutil.NewImageFromFile("res/graphics/" + s + ".png"))
	return

}

func GetSqr(n int, off int, s string) *ebiten.Image {
	var a = getPng(s)
	x := n % off * 16
	y := n / off * 16

	return a.SubImage(image.Rect(x, y, x+16, y+16)).(*ebiten.Image)
}

func GetLine(s string, l int, h int) (d []ebiten.Image) {

	for i := 0; i < l; i++ {

		d[l] = *GetSqr(h*4+i, 4, s)
	}
	return
}

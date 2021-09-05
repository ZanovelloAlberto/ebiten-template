package res

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Player []ebiten.Image
)

func init() {

	img, _, _ := ebitenutil.NewImageFromFile("res/graphics/player.png")
	Player = pack(*img, 3)

}

func pack(img ebiten.Image, l int) (d []ebiten.Image) {
	for i := 0; i < l; i++ {
		// u := image.Rect(i*16, 0, i*16+16, 16)
		// d[i] = img.

	}
	return
}

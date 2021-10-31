package assets

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed basictiles.png
	Basictiles []byte

	// go:embed characters.png
	Characters []byte
)

const (
	size   = 16
	assets = "/assets/"
	char   = assets + "characters.png"
	tiles  = assets + "basictiles.png"
)

func init() {

}

// func GetTile(x int, y int) (r *ebiten.Image) {
// 	// imgs, _, _ := ebitenutil.NewImageFromFile("res/basictiles.png")
// 	xv := x * size
// 	yv := y * size
// 	// rect := image.Rect(xv, yv, xv+size, yv+size)
// 	r = GetSqr(xv, yv, tiles)

// 	return

// }
// func GetCharacter(n int) (r [12]ebiten.Image) {
// 	vy := (n / 4) * 16 * 4
// 	vx := (n % 4) * 48

// 	for i := 0; i < 12; i++ {
// 		r[i] = *GetSqr(vx+(i%3)*16, vy+(i/3)*16, char)

// 	}
// 	return
// }

// func GetSqr(x int, y int, res string) *ebiten.Image {

// 	imgs, _, _ := ebitenutil.NewImageFromFile(res)
// 	return imgs.SubImage(image.Rect(x, y, x+16, y+16)).(*ebiten.Image)
// }

func GetImage(val []byte) *ebiten.Image {

	img, _, err := image.Decode(bytes.NewReader(val))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)

}

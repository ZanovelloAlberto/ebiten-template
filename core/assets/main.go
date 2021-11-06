package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed basictiles.png
	Basictiles []byte

	//go:embed characters.png
	Characters []byte
)

const (
	size = 16
)

func init() {
	// println(string(Basictiles))
}

func GetTile(x int, y int) (r *ebiten.Image) {
	// imgs, _, _ := ebitenutil.NewImageFromFile("res/basictiles.png")
	xv := x * size
	yv := y * size
	// rect := image.Rect(xv, yv, xv+size, yv+size)
	r = GetSqr(&Basictiles, xv, yv)

	return

}

func GetCharacter(n int) (r [12]ebiten.Image) {
	vy := (n / 4) * 16 * 4
	vx := (n % 4) * 48

	for i := 0; i < 12; i++ {
		r[i] = *GetSqr(&Characters, vx+(i%3)*16, vy+(i/3)*16)
	}
	return
}

func GetSqr(data *[]byte, y int, x int) *ebiten.Image {

	return GetImage(data).SubImage(image.Rect(x, y, x+16, y+16)).(*ebiten.Image)
}

func GetImage(val *[]byte) *ebiten.Image {

	img, err := png.Decode(bytes.NewReader(*val))
	if err != nil {
		panic(err.Error() + " non va la decodifica")
	}
	return ebiten.NewImageFromImage(img)

}

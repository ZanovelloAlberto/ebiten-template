package assets

import (
	"bytes"
	_ "embed"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed 2048_icon.png
	Icon []byte
)

func GetImage(val *[]byte) *ebiten.Image {

	img, err := png.Decode(bytes.NewReader(*val))
	if err != nil {
		panic(err.Error() + " non va la decodifica")
	}
	return ebiten.NewImageFromImage(img)
}

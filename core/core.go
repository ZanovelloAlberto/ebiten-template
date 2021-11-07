package core

import (
	"image/color"

	S "github.com/ZanovelloAlberto/EbitenGame/core/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

// CONSTUCTOR
func CreateCore(screenWidth int, screenHeight int) *Core {
	return &Core{
		screenWidth:  screenHeight,
		screenHeight: screenHeight,
	}
}

type Core struct {
	ebiten.Game
	screenWidth  int
	screenHeight int
}

func (core *Core) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 200, 200, 0xff})

	S.PPlayer.Draw(screen)
	S.PLoading.Draw(screen)
}

func (uno *Core) Update() error {

	S.PLoading.Update()
	S.PPlayer.Update()
	return nil
}

func (g *Core) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Run(core *Core) {

	ebiten.SetMaxTPS(1)
	if err := ebiten.RunGame(core); err != nil {
		panic(err.Error())
	}
}

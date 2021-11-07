package main

import (
	"github.com/ZanovelloAlberto/EbitenGame/core"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 100
	screenHeight = 100
)

func main() {
	a := &Desktop{
		core: *core.CreateCore(screenWidth, screenHeight),
	}
	ebiten.SetMaxTPS(5)
	if err := ebiten.RunGame(a); err != nil {
		panic(err.Error())
	}

}

type Desktop struct {
	core core.Core
}

func (uno *Desktop) Draw(screen *ebiten.Image) {
	uno.core.Draw(screen)
}

func (uno *Desktop) Update() error {
	return uno.core.Update()
}

func (g *Desktop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.core.Layout(outsideWidth, outsideHeight)
}

package main

import (
	"log"

	"github.com/ZanovelloAlberto/EbitenGame/core/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

func main() {
	a := &Desktop{}
	ebiten.SetMaxTPS(15)
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

type Desktop struct {
	core game.Game
}

func (uno *Desktop) Draw(screen *ebiten.Image) {
	uno.core.Draw(screen)

}

func (uno *Desktop) Update() error {
	return uno.Update()
}

func (g *Desktop) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

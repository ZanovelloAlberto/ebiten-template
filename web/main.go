package main

import (
	"log"

	twenty48 "github.com/ZanovelloAlberto/EbitenGame/2048"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(twenty48.ScreenWidth, twenty48.ScreenHeight)
	ebiten.SetWindowTitle("2048 (Ebiten Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

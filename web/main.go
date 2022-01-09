package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	game, err := NewGame()
	if err != nil {

	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(game); err != nil {
	}
}

package game

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	Vloading.Update()
	MainChar.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	// op.GeoM.Translate(screenWidth/2, screenHeight/2)
	// i := (g.count / 5) % frameNum
	// sx, sy := frameOX+i*frameWidth, frameOY
	// screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	screen.Fill(color.RGBA{100, 200, 200, 0xff})

	MainChar.Draw(screen)
	Vloading.Draw(screen)
}

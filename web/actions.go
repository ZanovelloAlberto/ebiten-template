package main

import (
	"fmt"
	"image/color"

	"github.com/ZanovelloAlberto/EbitenGame/core"
	"github.com/ZanovelloAlberto/EbitenGame/core/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
	"github.com/yohamta/furex/examples/shared"
)

type Actions struct {
	RootFlex      *furex.Flex
	isInitialized bool
	scoreLabel    *core.TextField
}

func NewActions() (*Actions, error) {
	actions := &Actions{
		isInitialized: false,
	}

	return actions, nil
}

func (a *Actions) Update() error {
	if a.isInitialized == false {
		a.buildUI()
		a.isInitialized = true
	}
	a.scoreLabel.Text = fmt.Sprint(core.Score)
	a.RootFlex.Update()

	return nil
}

func (a *Actions) Draw(screen *ebiten.Image) {
	a.RootFlex.Draw(screen)
}
func (a *Actions) buildUI() {
	// root container
	a.RootFlex = furex.NewFlex(ScreenWidth, ScreenHeight)
	a.RootFlex.Direction = furex.Column
	a.RootFlex.Justify = furex.JustifySpaceBetween
	a.RootFlex.AlignContent = furex.AlignContentCenter

	// top container
	top := furex.NewFlex(ScreenWidth-80, 100)
	top.SetMargin([]int{10, 10, 10, 10})
	top.Direction = furex.Row
	top.Justify = furex.JustifySpaceBetween
	top.AlignItems = furex.AlignItemCenter
	top.AddChild(core.NewPicture(assets.Icon, 100, 100))

	a.scoreLabel = core.NewTextField(color.RGBA{0, 0xff, 0, 0xff}, "0")
	top.AddChild(a.scoreLabel)
	a.RootFlex.AddChildContainer(top)

	// bottom container
	bottom := furex.NewFlex(ScreenWidth*6/8, 70)
	bottom.Direction = furex.Row
	bottom.SetMargin([]int{20, 20, 20, 20})
	bottom.Justify = furex.JustifySpaceBetween
	// bottom.AlignItems = furex.AlignItemEnd
	bottom.AddChild(core.NewButton(100, 50, func() {
		board, _ := core.NewBoard(4)
		core.GameBoard = board
	}, "NEW GAME"))
	bottom.AddChild(core.NewButton(100, 50, func() {
		core.GameBoard.LastMoveBack()
	}, "BACK"))
	a.RootFlex.AddChildContainer(bottom)
}

func buttonWithMargin(w, h int, margin []int) *shared.Button {
	b := shared.NewButton(w, h)
	b.SetMargin(margin)
	return b
}

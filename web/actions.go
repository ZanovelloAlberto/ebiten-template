package main

import (
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
	// top.AddChild(core.NewTextField(0, 0, core.FrameColor))
	top.AddChild(shared.NewBox(50, 50, color.RGBA{0, 0xff, 0, 0xff}))
	a.RootFlex.AddChildContainer(top)

	// bottom container
	bottom := furex.NewFlex(ScreenWidth, 70)
	bottom.Direction = furex.Row
	bottom.Justify = furex.JustifyCenter
	bottom.AlignItems = furex.AlignItemEnd
	bottom.AddChild(core.NewButton(200, 50, func() {
		board, _ := core.NewBoard(4)
		core.GameBoard = board
	}, "NEW GAME"))
	a.RootFlex.AddChildContainer(bottom)
}

func buttonWithMargin(w, h int, margin []int) *shared.Button {
	b := shared.NewButton(w, h)
	b.SetMargin(margin)
	return b
}

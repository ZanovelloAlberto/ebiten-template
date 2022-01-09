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
	top := furex.NewFlex(ScreenWidth-20, 70)
	top.Direction = furex.Row
	top.Justify = furex.JustifySpaceBetween
	top.AlignItems = furex.AlignItemCenter
	// top.AddChild(NewTextField(100, 50, color.RGBA{0xff, 0, 0, 0xff}))
	top.AddChild(core.NewPicture(assets.Icon, 50, 50))
	// print(string(icon))
	// top.AddChild(shared.NewBox(100, 30, color.RGBA{0xff, 0xff, 0xff, 0xff}))
	top.AddChild(shared.NewBox(50, 50, color.RGBA{0, 0xff, 0, 0xff}))
	a.RootFlex.AddChildContainer(top)

	// center
	// rootFlex.AddChild(shared.NewButton(100, 50))

	// bottom container
	bottom := furex.NewFlex(ScreenWidth, 70)
	bottom.Direction = furex.Row
	bottom.Justify = furex.JustifyCenter
	bottom.AlignItems = furex.AlignItemEnd
	bottom.AddChild(buttonWithMargin(50, 30, []int{5, 5, 10, 5}))
	bottom.AddChild(buttonWithMargin(50, 30, []int{5, 5, 10, 5}))
	bottom.AddChild(buttonWithMargin(50, 30, []int{5, 5, 10, 5}))
	bottom.AddChild(buttonWithMargin(50, 30, []int{5, 5, 10, 5}))
	a.RootFlex.AddChildContainer(bottom)
}
func buttonWithMargin(w, h int, margin []int) *shared.Button {
	b := shared.NewButton(w, h)
	b.SetMargin(margin)
	return b
}

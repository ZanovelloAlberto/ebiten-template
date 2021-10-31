package main

import (
	"core/core/assets"
	_ "embed"
	"os/exec"
)

//go:embed core/assets/basictiles.png
var s []byte

func main() {

	// var a = assets.GetImage(s)
	println(string(assets.Characters))
}

const (
	assets_source_dir = "./"
)

func Embed() {

	hcp := exec.Command("cp", "-r", "assets/", assets_source_dir)
	hcp.Start()
}

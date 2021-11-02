package main

import "github.com/ZanovelloAlberto/EbitenGame/core/assets"

// "core/core/assets"

var (
// go:embed core/assets/basictiles.png
// s []byte
)

func main() {

	// var a = assets.GetImage(s)
	println(string(assets.Basictiles))
	// println(string(assets.Basicatiles))
}

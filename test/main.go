package main

import (
	_ "embed"
)

var (
	//go:embed bu.png
	img []byte
)

func main() {

	// a := bytes.NewBuffer(img)

	// img, e := png.Decode(a)
	// // println(e.Error())
	// //	image: unknown format

	// println(img.Bounds().Dx())
	//

	// println(e.Error())
	// (0x0,0x0)

	// println(string(img))
	// the text of the image seem a little different

}

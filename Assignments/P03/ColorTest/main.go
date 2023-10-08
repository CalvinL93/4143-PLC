package main

import (
	"github.com/CalvinL93/img_mod/Colours"
	"github.com/CalvinL93/img_mod/GetPic"
	"github.com/CalvinL93/img_mod/Grayscale"
	"github.com/CalvinL93/img_mod/Text"
)

func main() {
	GetPic.GetPic()
	Colours.Colours()
	Grayscale.Grayscale()
	Text.Text()
}

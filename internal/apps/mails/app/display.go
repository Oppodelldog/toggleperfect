package app

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func CreateDisplayImage() image.Image {
	const screenW = 264
	dc := gg.NewContext(screenW, 176)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)
	headline := "***** MAILS PERFECT *****"
	headlineW, _ := dc.MeasureString(headline)
	dc.DrawString("***** MAILS PERFECT *****", screenW/2-headlineW/2, 24)

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

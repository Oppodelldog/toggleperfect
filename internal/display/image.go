package display

import (
	"image"
	"image/color"
)

func convertToGray(img image.Image) *image.Gray {
	b := img.Bounds()
	grayImage := image.NewGray(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := b.Max.X; x > 0; x-- {
			grayImage.Set(b.Max.X-x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return grayImage
}

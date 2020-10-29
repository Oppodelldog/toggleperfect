package app

import (
	"fmt"
	"github.com/MaxHalford/halfgone"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
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
	birthday := loadNextBirthday(time.Now(), findBirthdayFiles())

	dc.SetFontFace(face)
	headline := fmt.Sprintf("%s in %v days", birthday.Name, int(time.Until(birthday.GetNextDate(time.Now())).Hours()/24))
	headlineW, _ := dc.MeasureString(headline)
	dc.DrawString(headline, screenW/2-headlineW/2, 24)

	f, err := os.Open(birthday.FilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	birthdayChild, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	birthdayChildResized := resize.Resize(0, 130, birthdayChild, resize.Lanczos3)
	ditheredImage := halfgone.ThresholdDitherer{Threshold: 100}.Apply(display.ConvertToGray(birthdayChildResized))

	dc.DrawImage(ditheredImage, screenW/2-ditheredImage.Bounds().Dx()/2, 40)
	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

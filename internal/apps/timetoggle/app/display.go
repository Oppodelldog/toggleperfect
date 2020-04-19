package app

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"time"
)

func CreateDisplayImage() image.Image {
	data := getTimeData()

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
	headline := "***** TOGGLE PERFECT *****"
	headlineW, _ := dc.MeasureString(headline)
	dc.DrawString("***** TOGGLE PERFECT *****", screenW/2-headlineW/2, 24)

	dateToday := time.Now().Format("Monday, 02.01.2006")
	dateTodayW, _ := dc.MeasureString(dateToday)
	dc.DrawString(dateToday, screenW/2-dateTodayW/2, 50)

	timeNow := time.Now().Format("15:04")
	timeNowW, _ := dc.MeasureString(timeNow)
	dc.DrawString(timeNow, screenW/2-timeNowW/2, 76)

	face = truetype.NewFace(font, &truetype.Options{
		Size: 12,
	})
	dc.SetFontFace(face)
	dc.DrawString("Time to work today:", 10, 110)
	dc.DrawString("Time to work this week:", 10, 150)

	face = truetype.NewFace(font, &truetype.Options{
		Size: 24,
	})
	dc.SetFontFace(face)
	dc.DrawString((data.TimeToWorkToday - data.TimeWorkedToday).String(), 148, 110)
	dc.DrawString(data.RemainingTimeThisWeek.String(), 148, 150)

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

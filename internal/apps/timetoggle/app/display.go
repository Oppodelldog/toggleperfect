package app

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
)

const screenW = 264
const screenH = 176

func getRegularFont() *truetype.Font {
	// https://de.wikipedia.org/wiki/Windows_Glyph_List_4
	fnt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}

	return fnt
}

func getBoldFont() *truetype.Font {
	// https://de.wikipedia.org/wiki/Windows_Glyph_List_4
	fnt, err := truetype.Parse(gobold.TTF)
	if err != nil {
		panic("")
	}

	return fnt
}

func drawHeadline(dc *gg.Context) {
	drawHCentered(dc, 20, "══════╡ toggle perfect ╞══════")
}

type pos struct {
	x float64
	y float64
}

func drawButton(dc *gg.Context, face font.Face, fnt *truetype.Font, b button) {
	dc.SetRGB(0, 0, 0)
	dc.DrawCircle(b.position.x, b.position.y, b.radius)
	dc.Fill()
	face = truetype.NewFace(fnt, &truetype.Options{
		Size: 20,
	})
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.DrawString(b.symbol, b.position.x+b.symbolPos.x, b.position.y+b.symbolPos.y)

	face = truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)
	dc.SetRGB(0, 0, 0)
	dc.DrawString(b.caption, b.captionPos.x, b.captionPos.y+4)
}

func drawTodayCapture(dc *gg.Context, fnt *truetype.Font, i float64, s string) {
	dc.SetFontFace(truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	}))

	drawHCentered(dc, i, s)
}

func drawProjectDescription(dc *gg.Context, fnt *truetype.Font, y float64, description string) {
	dc.SetFontFace(truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	}))

	drawHCentered(dc, y, description)
}

func drawProjectName(dc *gg.Context, fnt *truetype.Font, y float64, name string) {
	dc.SetFontFace(truetype.NewFace(fnt, &truetype.Options{
		Size: 20,
	}))
	drawHCentered(dc, y, name)
}

func drawHCentered(dc *gg.Context, y float64, text string) {
	textWidth, _ := dc.MeasureString(text)
	dc.DrawString(text, screenW/2-textWidth/2, y)
}

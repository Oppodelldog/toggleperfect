package app

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"time"
)

const screenW = 264
const screenH = 176

func CreateStartScreen(projects []Project) image.Image {

	dc := gg.NewContext(screenW, screenH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fnt := getFont()
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)
	drawHeadline(dc)

	if len(projects) > 0 {
		drawButtonHint(dc, face, fnt, 25, 50, "▼", "")

		face := truetype.NewFace(fnt, &truetype.Options{
			Size: 16,
		})
		dc.SetFontFace(face)

		const vMargin = 20
		const xOffset = 50.0
		const yOffset = 50.0

		var maxWidth = 0.0

		y := yOffset
		for _, project := range projects {
			text := project.Name
			textWidth, _ := dc.MeasureString(text)
			if maxWidth < textWidth {
				maxWidth = textWidth
			}
			dc.DrawString(text, xOffset, y)
			y += vMargin
		}
		y = yOffset
		for _, project := range projects {
			text := project.Capture
			dc.DrawString(text, xOffset+maxWidth+12, y)
			y += vMargin
		}
	} else {
		drawHCentered(dc, 80, "No Projects available")
	}

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

func getFont() *truetype.Font {

	// https://de.wikipedia.org/wiki/Windows_Glyph_List_4

	fnt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	return fnt

}

func drawHeadline(dc *gg.Context) {
	drawHCentered(dc, 20, "*** TIME TOGGLE PROJECTS ***")
}

func drawButtonHint(dc *gg.Context, face font.Face, fnt *truetype.Font, x, y float64, buttonText, text string) {
	dc.SetRGB(0, 0, 0)
	dc.DrawCircle(x+0, y+0, 14)
	dc.Fill()
	face = truetype.NewFace(fnt, &truetype.Options{
		Size: 20,
	})
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.DrawString(buttonText, x-8, y+6)

	face = truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)
	dc.SetRGB(0, 0, 0)
	dc.DrawString(text, x+20, y+4)
}

type Project struct {
	Name        string
	Description string
	Capture     string
}

func (p Project) startCapture() {
	err := repo.AddStart(repo.Capture{
		ID:        p.Name,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
	errStop := repo.AddStop(repo.Capture{
		ID:        p.Name,
		Timestamp: time.Now().Unix(),
	})
	if errStop != nil {
		panic(errStop)
	}
}

func (p Project) stopCapture() {
	err := repo.SetLatestStop(repo.Capture{
		ID:        p.Name,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
}

func CreateProjectScreen(p Project) image.Image {
	dc := gg.NewContext(screenW, screenH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fnt := getFont()
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)

	drawButtonHint(dc, face, fnt, 25, 20, "▲", "")
	drawButtonHint(dc, face, fnt, 25, screenH-54, "►", "")
	drawButtonHint(dc, face, fnt, 25, screenH-20, "◄", "")

	drawProjectName(dc, fnt, 60, p.Name)
	drawProjectDescription(dc, fnt, 90, p.Description)
	drawTodayCapture(dc, fnt, 128, p.Capture)

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
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

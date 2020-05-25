package app

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func CreateStartScreen(projectSummary ProjectSummary) image.Image {
	dc := gg.NewContext(screenW, screenH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fnt := getRegularFont()
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: 18,
	})
	dc.SetFontFace(face)
	drawHeadline(dc)

	if len(projectSummary.Projects) > 0 {
		drawButton(dc, face, fnt, buttonDown().SetY(20).SetX(20))
		drawButton(dc, face, fnt, buttonPreviousDay().SetY(54).SetX(20))
		drawButton(dc, face, fnt, buttonPagination().SetY(88).SetX(20))
		face = truetype.NewFace(fnt, &truetype.Options{
			Size: 12,
		})
		dc.SetFontFace(face)
		drawHCentered(dc, 42, projectSummary.Date.Format("Mon, 02 Jan"))

		face := truetype.NewFace(fnt, &truetype.Options{
			Size: 16,
		})
		dc.SetFontFace(face)

		const vMargin = 20
		const xOffset = 50.0
		const yOffset = 68.0

		var maxWidth = 0.0

		y := yOffset
		for _, item := range projectSummary.Pagination.GetCurrentPageItems(projectSummary.Projects) {
			project := item.(Project)
			if project.Capture == "" {
				continue
			}
			text := project.Name
			textWidth, _ := dc.MeasureString(text)
			if maxWidth < textWidth {
				maxWidth = textWidth
			}
			dc.DrawString(text, xOffset, y)
			y += vMargin
		}
		y = yOffset
		for _, item := range projectSummary.Pagination.GetCurrentPageItems(projectSummary.Projects) {
			project := item.(Project)
			text := project.Capture
			dc.DrawString(text, xOffset+maxWidth+12, y)
			y += vMargin
		}

		drawHCentered(dc, y, fmt.Sprintf("%v/%v", projectSummary.Pagination.Page, projectSummary.Pagination.GetLastPage()))
	} else {
		drawHCentered(dc, 80, "No Projects available")
	}

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

func CreateProjectScreen(p Project) image.Image {
	dc := gg.NewContext(screenW, screenH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fnt := getRegularFont()
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: 14,
	})
	dc.SetFontFace(face)

	drawHeadline(dc)

	drawButton(dc, face, fnt, buttonUp().SetX(20).SetY(20))
	drawButton(dc, face, getBoldFont(), buttonClose().SetX(20).SetY(54))

	drawButton(dc, face, fnt, buttonRight().SetY(screenH-54).SetX(20))
	drawButton(dc, face, fnt, buttonLeft().SetY(screenH-20).SetX(20))

	drawProjectName(dc, fnt, 60, p.Name)
	drawProjectDescription(dc, fnt, 90, p.Description)
	drawTodayCapture(dc, fnt, 128, p.Capture)

	if p.Closed {
		dc.SetFontFace(truetype.NewFace(fnt, &truetype.Options{
			Size: 44,
		}))
		drawHCentered(dc, 150, "√")
	}

	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

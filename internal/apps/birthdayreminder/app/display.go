package app

import (
	"github.com/MaxHalford/halfgone"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/log"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"image"
	"image/png"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
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
	dc.SetFontFace(face)
	headline := "***** LED DEMO *****"
	headlineW, _ := dc.MeasureString(headline)
	dc.DrawString("*** BIRTHDAY ***", screenW/2-headlineW/2, 24)

	fileName := loadNextBirthday(time.Now(), findBirthdayFiles())
	f, err := os.Open(path.Join(getBirthdayDir(), fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	birthdayChild, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	birthdayChildResized := resize.Resize(0, 130, birthdayChild, resize.Lanczos3)
	ditheredImage := halfgone.ThresholdDitherer{Threshold: 60}.Apply(display.ConvertToGray(birthdayChildResized))

	dc.DrawImage(ditheredImage, screenW/2-ditheredImage.Bounds().Dx()/2, 40)
	dc.InvertX()
	dc.Push()
	dc.InvertY()
	dc.Push()

	return dc.Image()
}

func findBirthdayFiles() []string {
	var fileNames []string
	birthdayDir := getBirthdayDir()
	err := filepath.Walk(birthdayDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".png" {
			fileNames = append(fileNames, info.Name())
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return fileNames
}

func getBirthdayDir() string {
	birthdayDir, hassBirthdayDir := os.LookupEnv("TOGGLE_PERFECT_BIRTHDAY_DIR")

	if !hassBirthdayDir {
		birthdayDir = path.Dir(util.GetExecutableDir())
	}
	return birthdayDir
}

func loadNextBirthday(today time.Time, birthdayFiles []string) string {
	birthdayEntries := map[string]int64{}

	for _, fileName := range birthdayFiles {
		log.Print(fileName)
		name := fileName[:strings.Index(fileName, ".png")]
		parts := strings.SplitN(name, "_", 2)
		log.Printf("%#v", parts)
		thisYear := today.Year()
		dateParts := strings.SplitN(parts[1], "-", 2)
		birthdate := time.Date(thisYear, time.Month(mustInt(dateParts[1])), mustInt(dateParts[0]), 12, 0, 0, 0, time.UTC)

		if birthdate.Unix() < today.Unix() {
			birthdate = time.Date(thisYear+1, time.Month(mustInt(dateParts[1])), mustInt(dateParts[0]), 12, 0, 0, 0, time.UTC)
		}

		timeToBirthday := birthdate.Unix() - today.Unix()
		birthdayEntries[fileName] = timeToBirthday
	}

	nearest := int64(math.MaxInt64)
	nearestName := ""

	for name, t := range birthdayEntries {
		if t > 0 && t < nearest {
			nearest = t
			nearestName = name
		}
	}

	return nearestName
}

func mustInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

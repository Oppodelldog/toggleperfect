package app

import (
	"errors"
	"fmt"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidFormat = errors.New("invalid birthday format")

type Birthday struct {
	Filename string
	FilePath string
	Name     string
	Day      int
	Month    int
}

func (b Birthday) GetNextDate(refDate time.Time) time.Time {
	birthdate := time.Date(refDate.Year(), time.Month(b.Month), b.Day, 12, 0, 0, 0, time.UTC)
	if birthdate.Unix() < refDate.Unix() {
		birthdate = time.Date(refDate.Year()+1, time.Month(b.Month), b.Day, 12, 0, 0, 0, time.UTC)
	}

	return birthdate
}

func NewBirthday(filePath string) (Birthday, error) {
	fileName := path.Base(filePath)
	ext := filepath.Ext(fileName)

	if ext != ".png" {
		return Birthday{}, fmt.Errorf("%w: must be .png, but got %s", ErrInvalidFormat, ext)
	}

	extPos := strings.Index(fileName, ext)
	name := fileName[:extPos]

	parts := strings.SplitN(name, "_", 2)
	if len(parts) != 2 {
		return Birthday{}, fmt.Errorf("%w: name_date must be two parts, found %v", ErrInvalidFormat, len(parts))
	}

	dateParts := strings.SplitN(parts[1], "-", 2)
	if len(dateParts) != 2 {
		return Birthday{}, fmt.Errorf("%w: day-month must be two parts, found %v", ErrInvalidFormat, len(dateParts))
	}

	return Birthday{
		Filename: fileName,
		FilePath: filePath,
		Name:     parts[0],
		Day:      mustInt(dateParts[0]),
		Month:    mustInt(dateParts[1]),
	}, nil
}

func findBirthdayFiles() []string {
	var fileNames []string
	birthdayDir := getBirthdayDir()
	err := filepath.Walk(birthdayDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".png" {
			fileNames = append(fileNames, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return fileNames
}

func getBirthdayDir() string {
	birthdayDir, hasBirthdayDir := os.LookupEnv("TOGGLE_PERFECT_BIRTHDAY_DIR")

	if !hasBirthdayDir {
		birthdayDir = path.Dir(util.GetExecutableDir())
	}
	return birthdayDir
}

func loadNextBirthday(today time.Time, birthdayFiles []string) Birthday {
	birthdayEntries := map[Birthday]int64{}

	for _, fileName := range birthdayFiles {
		birthDay, err := NewBirthday(fileName)
		if err != nil {
			logrus.Warnf("could not load birthday '%s': %v", fileName, err)
			continue
		}

		birthdate := birthDay.GetNextDate(today).Unix()

		timeToBirthday := birthdate - today.Unix()
		birthdayEntries[birthDay] = timeToBirthday
	}

	nearest := int64(math.MaxInt64)
	var nearestBirthday Birthday

	for name, t := range birthdayEntries {
		if t > 0 && t < nearest {
			nearest = t
			nearestBirthday = name
		}
	}

	return nearestBirthday
}

func mustInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Capture struct {
	ID        string
	Timestamp int64
}

type CaptureFile struct {
	ID     string
	Starts []int64
	Stops  []int64
}

func GetAllCaptures() ([]CaptureFile, error) {
	var monthDirs []string
	err := filepath.Walk(capturesDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != capturesDir {
			monthDirs = append(monthDirs, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var files []*os.File
	defer func() {
		for _, f := range files {
			err := f.Close()
			if err != nil {
				log.Printf("error in error cleanup while closing file: %v", err)
			}
		}
	}()

	for _, monthDir := range monthDirs {
		monthFiles, err := getStorageFiles(monthDir, openCaptureFileForReadingFromDir)
		if err != nil {
			return nil, err
		}

		files = append(files, monthFiles...)
	}

	var captures []CaptureFile
	for _, f := range files {
		var capture CaptureFile
		err = json.NewDecoder(f).Decode(&capture)
		if err != nil {
			return nil, err
		}

		captures = append(captures, capture)
	}

	return captures, nil
}

func SetLatestStop(stop Capture) error {
	return mutateCaptureFile(stop.ID, func(captureFile CaptureFile) CaptureFile {
		if len(captureFile.Stops) == 0 {
			captureFile.Stops = append(captureFile.Stops, stop.Timestamp)
		} else {
			captureFile.Stops[len(captureFile.Stops)-1] = stop.Timestamp
		}
		captureFile.ID = stop.ID

		return captureFile
	})
}

func AddStart(start Capture) error {
	return mutateCaptureFile(start.ID, func(captureFile CaptureFile) CaptureFile {
		captureFile.Starts = append(captureFile.Starts, start.Timestamp)
		captureFile.ID = start.ID

		return captureFile
	})
}

func AddStop(stop Capture) error {
	return mutateCaptureFile(stop.ID, func(captureFile CaptureFile) CaptureFile {
		captureFile.Stops = append(captureFile.Stops, stop.Timestamp)
		captureFile.ID = stop.ID

		return captureFile
	})
}

type mutateCaptureFileFunc func(CaptureFile) CaptureFile

func mutateCaptureFile(ID string, f mutateCaptureFileFunc) error {
	captureFile, err := openCaptureForCurrentMonth(ID)
	if err != nil {
		return err
	}

	captureFile = f(captureFile)

	return saveCaptureForCurrentMonth(captureFile)
}

func openCaptureForCurrentMonth(ID string) (CaptureFile, error) {
	err := ensureFileForCurrentMonth(ID)
	if err != nil {
		return CaptureFile{}, fmt.Errorf("error ensuring file: %v", err)
	}
	f, err := openCaptureFileForReadingForCurrentMonth("", ID)
	if err != nil {
		return CaptureFile{}, fmt.Errorf("error opening file: %v", err)
	}
	defer closeWithPanic(f)

	var captureFile CaptureFile
	err = json.NewDecoder(f).Decode(&captureFile)
	if err != nil && err != io.EOF {
		return CaptureFile{}, fmt.Errorf("error decoding file: %v", err)
	}

	return captureFile, nil
}

func ensureFileForCurrentMonth(ID string) error {
	fileName := getFileNameForCurrentMonth("", ID)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return saveCaptureForCurrentMonth(CaptureFile{ID: ID})
	}

	return nil
}

func saveCaptureForCurrentMonth(captureFile CaptureFile) error {
	f, err := openCaptureFileForWritingForCurrentMonth("", captureFile.ID)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %v", err)
	}
	defer closeWithPanic(f)

	return json.NewEncoder(f).Encode(captureFile)
}

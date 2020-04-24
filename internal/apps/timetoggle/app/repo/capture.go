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

func GetCaptures() ([]CaptureFile, error) {
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

func AddStart(start Capture) error {
	f, err := openCaptureFileForReadingAndWritingForCurrentMonth("", start.ID)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer closeWithPanic(f)

	var captureFile CaptureFile
	err = json.NewDecoder(f).Decode(&captureFile)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error decoding file: %v", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error seeking: %v", err)
	}
	captureFile.Starts = append(captureFile.Starts, start.Timestamp)
	captureFile.ID = start.ID

	return json.NewEncoder(f).Encode(captureFile)
}

func AddStop(stop Capture) error {
	f, err := openCaptureFileForReadingAndWritingForCurrentMonth("", stop.ID)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer closeWithPanic(f)

	var captureFile CaptureFile
	err = json.NewDecoder(f).Decode(&captureFile)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error decoding file: %v", err)
	}
	_, err = f.Seek(0, 0)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error seeking: %v", err)
	}

	captureFile.Stops = append(captureFile.Stops, stop.Timestamp)
	captureFile.ID = stop.ID

	return json.NewEncoder(f).Encode(captureFile)
}

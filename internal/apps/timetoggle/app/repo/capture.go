package repo

import (
	"encoding/json"
	"fmt"
	"io"
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

func AddStart(start Capture) error {
	f, err := openCaptureFileForWriting(start.ID)
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
	f, err := openCaptureFileForWriting(stop.ID)
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

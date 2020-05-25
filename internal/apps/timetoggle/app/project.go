package app

import (
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app/repo"
)

type ProjectSummary struct {
	Date       time.Time
	Projects   []Project
	Pagination Pagination
}

type Project struct {
	Name        string
	Description string
	Closed      bool
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

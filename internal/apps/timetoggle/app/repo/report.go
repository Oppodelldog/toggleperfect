package repo

import (
	"time"
)

type ReportCapturesList struct {
	Projects []ReportCapturesCapture
}

type ReportCapturesCapture struct {
	ID                  string
	NumberOfTimesWorked int64
	TimeWorked          int64
	TimeWorkedDisplay   string
}

func GetTodayCaptures() (ReportCapturesList, error) {
	now := time.Now()
	minTime := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.UTC)
	maxTime := minTime.Add(time.Hour * 16).Add(time.Nanosecond * -1)

	return GetReportCaptures(minTime, maxTime)
}
func GetMonthCaptures() (ReportCapturesList, error) {
	now := time.Now()
	minTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	maxTime := minTime.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	return GetReportCaptures(minTime, maxTime)
}

func GetReportCaptures(minTime time.Time, maxTime time.Time) (ReportCapturesList, error) {
	captures, err := GetCaptures()
	if err != nil {
		return ReportCapturesList{}, err
	}

	result := ReportCapturesList{}
	for _, c := range captures {
		var secondsWorked int
		var numberOfTimesWorked int64
		for i, start := range c.Starts {
			if minTime.Unix() > start || maxTime.Unix() < start {
				continue
			}

			if len(c.Stops) > i {
				secondsWorked += int(c.Stops[i] - start)
				numberOfTimesWorked++
			}
		}

		if numberOfTimesWorked > 0 {
			result.Projects = append(result.Projects, ReportCapturesCapture{
				ID:                  c.ID,
				TimeWorked:          int64(secondsWorked),
				TimeWorkedDisplay:   (time.Duration(secondsWorked) * time.Second).String(),
				NumberOfTimesWorked: numberOfTimesWorked,
			})
		}
	}

	return result, nil
}

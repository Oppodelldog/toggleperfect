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

func TimeSpanMonth(t time.Time) TimeSpan {
	minTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	maxTime := minTime.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	return TimeSpan{From: minTime, To: maxTime}
}

func TimeSpanDay(t time.Time) TimeSpan {
	minTime := time.Date(t.Year(), t.Month(), t.Day(), 6, 0, 0, 0, time.UTC)
	maxTime := minTime.Add(time.Hour * 16).Add(time.Nanosecond * -1)

	return TimeSpan{From: minTime, To: maxTime}
}

type TimeSpan struct {
	From time.Time
	To   time.Time
}

func GetTodayCaptures() (ReportCapturesList, error) {
	return GetReportCaptures(TimeSpanDay(time.Now()))
}

func GetMonthCaptures() (ReportCapturesList, error) {
	return GetReportCaptures(TimeSpanMonth(time.Now()))
}

func GetReportCaptures(timespan TimeSpan) (ReportCapturesList, error) {
	captures, err := GetAllCaptures()
	if err != nil {
		return ReportCapturesList{}, err
	}

	result := ReportCapturesList{}
	for _, c := range captures {
		var secondsWorked int
		var numberOfTimesWorked int64
		for i, start := range c.Starts {
			if timespan.From.Unix() > start || timespan.To.Unix() < start {
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

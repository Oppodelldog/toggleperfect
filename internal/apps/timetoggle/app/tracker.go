package app

import "time"

type State int

const (
	Done State = iota
	Work
	Pause
)

var state = Done
var startTime time.Time
var timeWorkedToday time.Duration
var timePauseToday time.Duration
var timeWorkedThisWeek = time.Hour*8*2 + time.Minute*27

func StartWork() {
	state = Work
	startTime = time.Now()
}

func StopWork() {
	state = Done
	timeWorkedToday += time.Since(startTime)
}

func StartPause() {
	state = Pause
	timeWorkedToday += time.Since(startTime)
	startTime = time.Now()
}

func StopPause() {
	state = Work
	timePauseToday = time.Since(startTime)
}

func TimeWorkedToday() time.Duration {
	return timeWorkedToday
}

func TimePausedToday() time.Duration {
	return timePauseToday
}

func TimeWorkedWekk() time.Duration {
	return timeWorkedThisWeek + timeWorkedToday
}

func CurrentState() State {
	return state
}

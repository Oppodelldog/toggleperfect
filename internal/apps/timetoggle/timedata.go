package timetoggle

import "time"

type Data struct {
	TimeWorkedToday       time.Duration
	TimeToWorkToday       time.Duration
	RemainingTimeThisWeek time.Duration
}

func getTimeData() Data {

	return Data{
		TimeWorkedToday:       time.Second * 0,
		TimeToWorkToday:       time.Hour * 4,
		RemainingTimeThisWeek: time.Hour * 40,
	}
}

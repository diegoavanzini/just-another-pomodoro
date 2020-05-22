package models

import "time"

type CurrentPomodoro struct {
	TimeDuration      int
	PauseDuration     int
	CurrentTimerValue string
}

type PomodoroPosition struct {
	TimeStarted     time.Time
	TimeDuration    int
	CurrentPosition int
}
